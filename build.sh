STAGING=develop
VERSION="v1"
export BUILD_TIME=$(date)
export BUILD_COMMIT="$(git rev-parse --short HEAD)"
export PROJECT_DIR=$(pwd)

set -a
  [ -f .env ] && . .env
set +a

export VERSION=$VERSION
export PROJECT_DIR=$PROJECT_DIR

# Compile output dir
BUILD_DIR="$PROJECT_DIR/build"
APP_NAME="project-v-go-noti"
# Print information before run build
echo "===== Building script ====="
echo "Staging: $STAGING"
echo "build_dir: $BUILD_DIR"
echo "Name: $APP_NAME"
echo "Version: $VERSION"
echo "Build at: $BUILD_TIME"
echo "Build commit: $BUILD_COMMIT"
echo "============================"


echo "[INF] Build main server"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -X 'main.ProgramName=$APP_NAME' -X 'main.Version=$VERSION' -X 'main.BuildTime=$BUILD_TIME'" -o "$BUILD_DIR/server" ./cmd/admin/main.go
cp -pr $PROJECT_DIR/blockchain/ $BUILD_DIR/blockchain/
echo ">>> Building docker image"
docker build -t "project-v-go-noti:v0.0.6" .
