#!/bin/bash

flash()
{
   local what="$1"
   if ! lsblk | grep -q sdc ; then 
      echo "*** sdc is missing"
      exit 1
   fi
   echo flashing
   sudo -s -- << EOF
echo "Mount device"
mount /dev/sdc /mnt || exit $?

echo "Install $what firmware"
cp build/${what}/zephyr/zmk.uf2 /mnt

echo "Unmount"
sync
umount /mnt
EOF
}

if [ -z "$1" ]; then
   mkdir -p build/{left,right}
   west build -d build/left -b nice_nano -- -DSHIELD=ergo_s1_oe_left || exit $?
   west build -d build/right -b nice_nano -- -DSHIELD=ergo_s1_oe_right || exit $?
elif [ "$1" == "l" ]; then
   mkdir -p build/left
   west build -d build/left -b nice_nano -- -DSHIELD=ergo_s1_oe_left || exit $?
elif [ "$1" == "r" ]; then
   mkdir -p build/right
   west build -d build/right -b nice_nano -- -DSHIELD=ergo_s1_oe_right || exit $?
elif [ "$1" == "ll" ]; then
   flash "left"
elif [ "$1" == "rr" ]; then
   flash "right"
else
   echo "*** Unknown argument!" >&2
   exit 1
fi

