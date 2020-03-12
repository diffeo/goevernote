#/bin/sh
#
# from: https://github.com/evernote/evernote-thrift/issues/10#issuecomment-324966201

set -ex

PARAMS="-out . -gen go:package_prefix=github.com/diffeo/goevernote/,thrift_import=github.com/apache/thrift/lib/go/thrift"
THRIFT_REPO="../evernote-thrift"
THRIFT_FILES_DIR="$THRIFT_REPO/src"

if [ ! -d "$THRIFT_REPO" ]; then
    git clone git@github.com:evernote/evernote-thrift.git "$THRIFT_REPO"
fi

thrift $PARAMS $THRIFT_FILES_DIR/Errors.thrift
thrift $PARAMS $THRIFT_FILES_DIR/Limits.thrift
thrift $PARAMS $THRIFT_FILES_DIR/NoteStore.thrift
thrift $PARAMS $THRIFT_FILES_DIR/Types.thrift
thrift $PARAMS $THRIFT_FILES_DIR/UserStore.thrift
