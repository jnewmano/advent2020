#!/bin/bash

set -eux

DAY=day$1

mkdir $DAY
cp parta.go $DAY/
cp runa $DAY/
cp runb $DAY/
