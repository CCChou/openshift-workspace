
if [ $# -ne 2 ];then
  echo deploy-nfs-sc.sh [server][share_path]
  exit 1
fi

SERVER=$1
SHARE_PATH=$2

curl -skSL https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/install-driver.sh | bash -s master --
cat <<EOF | oc apply -f -
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: nfs-csi
provisioner: nfs.csi.k8s.io
parameters:
  server: $SERVER
  share: $SHARE_PATH
reclaimPolicy: Retain
volumeBindingMode: Immediate
mountOptions:
  - hard
  - nfsvers=4.1
EOF
