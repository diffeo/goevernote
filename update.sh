#/bin/sh
#
# from: https://github.com/evernote/evernote-thrift/issues/10#issuecomment-324966201

set -ex

PARAMS="-out . -gen go:package_prefix=github.com/diffeo/goevernote/,thrift_import=github.com/apache/thrift/lib/go/thrift"
THRIFT_FILES_DIR="../evernote-thrift/src"

thrift $PARAMS $THRIFT_FILES_DIR/Errors.thrift
thrift $PARAMS $THRIFT_FILES_DIR/Limits.thrift
thrift $PARAMS $THRIFT_FILES_DIR/NoteStore.thrift
thrift $PARAMS $THRIFT_FILES_DIR/Types.thrift
thrift $PARAMS $THRIFT_FILES_DIR/UserStore.thrift
