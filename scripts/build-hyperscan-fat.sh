
CURDIR=$(pwd)

BUILDDIR=$CURDIR/vendor/hyperscan/hs_build

#FileType=None
#FileType=Debug
#FileType=MinSizeRel
#FileType=RelWithDebInfo
FileType=Release

ln -s $CURLDIR/vendor/boost vendor/hyperscan/include/boost

rm -fr $BUILDDIR && mkdir $BUILDDIR
cd $BUILDDIR && cmake ../                           \
                    -DCMAKE_BUILD_TYPE=$FileType    \
                    -DCMAKE_C_FLAGS="-fPIC "        \
                    -DCMAKE_CXX_FLAGS="-fPIC "      \
                    -DFAT_RUNTIME=1                 \
                && make -j 4                        \
                && cp lib/libhs.a ${CURDIR}         \
                && cd - || exit 1


