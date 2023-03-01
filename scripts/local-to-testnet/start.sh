#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

BUILDDIR=${SCRIPT_DIR}/../../build

mkdir -p $BUILDDIR

source ${SCRIPT_DIR}/config.sh

build_local_and_docker() {
   module="$1"
   folder="$2"
   title=$(printf "$module" | awk '{ print toupper($0) }')

   printf '%s' "Building $title Locally...  "
   cwd=$PWD

   cd $folder

#   GOBIN=$BUILDDIR go install -mod=readonly -trimpath -buildvcs=false ./... 2>&1 | grep -v -E "deprecated|keychain" | true
   echo $BUILDDIR go install -mod=readonly -trimpath -buildvcs=false ./... 2>&1 | grep -v -E "deprecated|keychain" | true
   local_build_succeeded=${PIPESTATUS[0]}

   cd $cwd

   if [[ "$local_build_succeeded" == "0" ]]; then
      echo "Done"
   else
      echo "Failed"
      return $local_build_succeeded
   fi

   echo "Building $title Docker...  "
   if [[ "$module" == "stayking" ]]; then
      image=Dockerfile
   else
      image=dockernet/dockerfiles/Dockerfile.$module
   fi
   echo "$image"
   DOCKER_BUILDKIT=1 docker build --tag soohoio:$module -f $image . | true
   docker_build_succeeded=${PIPESTATUS[0]}

   if [[ "$docker_build_succeeded" == "0" ]]; then
      echo "Done"
   else
      echo "Failed"
   fi
   return $docker_build_succeeded
}

ADMINS_FILE=${SCRIPT_DIR}/../../utils/admins.go
ADMINS_FILE_BACKUP=${SCRIPT_DIR}/../../utils/admins.go.main

replace_admin_address() {
   cp $ADMINS_FILE $ADMINS_FILE_BACKUP
   sed -i -E "s|sooho1k8c2m5cn322akk5wy8lpt87dd2f4yh9azg7jlh|$STAYKING_ADMIN_ADDRESS|g" $ADMINS_FILE
}

revert_admin_address() {
   mv $ADMINS_FILE_BACKUP $ADMINS_FILE
   rm -f ${ADMINS_FILE}-E
}

echo "Building STAYKING...";

replace_admin_address
if (build_local_and_docker stayking .) ; then
  revert_admin_address
else
  revert_admin_address
  exit 1
fi

echo "Building Gaia Testnet Relayer ...";
build_local_and_docker relayer deps/relayer


## Building done
echo "Building done"

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [[ $# -ne 0 && $1 = "i" ]]; then
  echo
  PS3="초기화 모드를 선택하셨습니다 계속 실행하시겠습니까?"
  COLUMNS=20
  options=(
    "Yes"
    "No"
  )
  select yn in "${options[@]}"; do
      case $yn in
          "Yes")    echo "chain nit mode";         break;;
          "No" )    exit;;
      esac
  done
# cleanup any stale state
  rm -rf $STATE $LOGS
  mkdir -p $STATE
  mkdir -p $LOGS
  # Start stayking chain
  echo "StayKing init mode..."
  bash ${SCRIPT_DIR}/init_stayking.sh $STAYKING_CHAIN_ID
fi

for chain_id in STAYKING; do
    num_nodes=$(GET_VAR_VALUE ${chain_id}_NUM_NODES)
    node_prefix=$(GET_VAR_VALUE ${chain_id}_NODE_PREFIX)

    log_file=$LOGS/${node_prefix}.log

    echo "Starting $chain_id chain"
    nodes_names=$(i=1; while [ $i -le $num_nodes ]; do printf "%s " ${node_prefix}${i}; i=$(($i + 1)); done;)

    $DOCKER_COMPOSE up -d $nodes_names
    $DOCKER_COMPOSE logs -f ${node_prefix}1 | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" > $log_file 2>&1 &
done

for chain_id in STAYKING; do
    printf "Waiting for $chain_id to start..."

    node_prefix=$(GET_VAR_VALUE ${chain_id}_NODE_PREFIX)
    log_file=$LOGS/${node_prefix}.log

    ( tail -f -n0 $log_file & ) | grep -q "finalizing commit of block"
    echo "Done"
done

sleep 5

if [[ $# -ne 0 && $1 = "i" ]]; then
  echo "add relayer keys and start relayers !"
  bash $SCRIPT_DIR/start_relayers.sh $1
else
  echo "start relayers !"
  bash $SCRIPT_DIR/start_relayers.sh
fi

echo "register host !"
bash $SCRIPT_DIR/register_host.sh

echo "create logs !"
$SCRIPT_DIR/create_logs.sh &

echo "Done"
