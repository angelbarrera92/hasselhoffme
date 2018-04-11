#!/bin/bash

################################
## Platform: Gnome            ##
################################


## Preparing full url of image ##
URL='https://media.cntraveller.in/wp-content/uploads/2016/08/davidhasselhofflead.jpg'

## Local directory where this script file placed
CURDIR="/tmp"

## Local directory for images
IMGDIR=$CURDIR"/david"

## Creating directory if not present ##
mkdir -m a=rwx -p $IMGDIR

## Local path of output image ##
LOCALIMG=$IMGDIR"/hasselhofflead.jpg"

## Download image to local path ##
curl $URL --output $LOCALIMG

## Setting desktop background image from local ##
gsettings set org.gnome.desktop.background picture-uri 'file://'$LOCALIMG
