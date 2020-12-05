#!/bin/bash

set -eux

DAY=day$1

mkdir $DAY
cp templates/* $DAY/
