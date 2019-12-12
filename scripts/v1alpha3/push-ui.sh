set -o errexit
set -o nounset
set -o pipefail

REGISTRY="quay.io/babylonhealth"
TAG="v1alpha3"
PREFIX="katib"
CMD_PREFIX="cmd"
MACHINE_ARCH=`uname -m`

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/../..

cd ${SCRIPT_ROOT}

usage() { echo "Usage: $0 [-t <tag>] [-r <registry>] [-p <prefix>]" 1>&2; exit 1; }

while getopts ":t::r::p:" opt; do
    case $opt in
        t)
            TAG=${OPTARG}
            ;;
        r)
            REGISTRY=${OPTARG}
            ;;
        p)
            PREFIX=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done
echo "Registry: ${REGISTRY}, tag: ${TAG}, prefix: ${PREFIX}"

echo "Building UI image..."
docker build -t ${REGISTRY}/${PREFIX}:katib-ui-${TAG} -f ${CMD_PREFIX}/ui/v1alpha3/Dockerfile .

echo "Pushing UI image..."
docker push ${REGISTRY}/${PREFIX}:katib-ui-${TAG}
