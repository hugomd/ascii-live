#!/bin/bash

# Generate ASCII art frames from a GIF file.

HOSTNAME=localhost:8124
IN_FILE=""
OUT_DIR="ansi_frames/"
NAME=""
DELAY=100
COLOR_DEPTH=8
WIDTH=64

# Get arguments
if [ $# -eq 0 ]; then
  echo "Usage: gen_frames.sh [-o output_dir] [-n name] [-d delay] [-c color_depth] [-w width] <gif_file>" >&2
  exit 1
fi

# Check dependencies
if ! command -v jp2a &> /dev/null; then
  echo "jp2a not found. Install it with: sudo apt install jp2a" >&2
  exit 1
fi
if ! command -v convert &> /dev/null; then
  echo "convert not found. Install it with: sudo apt install imagemagick" >&2
  exit 1
fi

# Get options
while getopts ":o:n:d:c:w:h" opt; do
  case $opt in
    o)
      OUT_DIR=$OPTARG
      ;;
    n)
      if [[ $OPTARG =~ ^[a-zA-Z0-9_]+$ ]]; then
        NAME=$OPTARG
      else
        echo "Invalid name: $OPTARG. Must be alphanumeric characters and underscores only." >&2
        exit 1
      fi
      ;;
    d)
      if [ $OPTARG -gt 0 ]; then
        DELAY=$OPTARG
      else
        echo "Invalid delay: $OPTARG. Must be an integer >0." >&2
        exit 1
      fi
      ;;
    c)
      if [ $OPTARG -eq 1 ] || [ $OPTARG -eq 8 ] || [ $OPTARG -eq 24 ]; then
        COLOR_DEPTH=$OPTARG
      else
        echo "Invalid color depth: $OPTARG. Must be 1 (monochrome), 8 (256 colors), or 24 (true color)." >&2
        exit 1
      fi
      ;;
    w)
      if [ $OPTARG -gt 0 ]; then
        WIDTH=$OPTARG
      else
        echo "Invalid width: $OPTARG. Must be an integer >0." >&2
        exit 1
      fi
      ;;
    h)
      echo "Usage: gen_frames.sh [options...] <gif_file>" >&1
      echo " -o <output_dir>    Output directory. Default: $OUT_DIR" >&1
      echo " -n <name>          Frame set name." >&1
      echo " -d <delay>         Frame delay in milliseconds. Default: $DELAY" >&1
      echo " -c <color_depth>   Color depth. Must be 1 (monochrome), 8 (256 colors), or 24 (true color). Default: $COLOR_DEPTH" >&1
      echo " -w <width>         Width of output. Default: $WIDTH" >&1
      exit 0
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
shift $((OPTIND -1))

# Check input file
if [ ! -f "$1" ]; then
  echo "File not found: $1" >&2
  exit 1
fi
if ! file "$1" | grep -q GIF; then
  echo "File is not a GIF: $1" >&2
  exit 1
fi
IN_FILE="$1"

# Get frame set name
if [ -z "$NAME" ] && ! [[ $IN_FILE =~ ^-.* ]]; then
  NAME=$(basename "$IN_FILE" .gif)
  if ! [[ $NAME =~ ^[a-zA-Z0-9_]+$ ]]; then
    NAME=$(echo "$NAME" | sed -E 's/[^a-zA-Z0-9_]+//g')
  fi
fi

# Set up output directory
mkdir -p "$OUT_DIR/$NAME"
rm -f "$OUT_DIR/$NAME/"*

# Generate PNG frames
convert -coalesce "$IN_FILE" "$OUT_DIR/$NAME/frame-%03d.png"

# Set delay
echo "$DELAY" > "$OUT_DIR/$NAME/delay.env"

# Generate ANSI frames
for f in "$OUT_DIR/$NAME/frame-"*.png; do
  frame=$(basename "$f" .png)
  jp2a --colors --color-depth=$COLOR_DEPTH --background=dark --width=$WIDTH "$f" > "$OUT_DIR/$NAME/$frame.ansi" && \
  rm "$f"
done

echo "[+] Done! Try it with:"
echo "Restart the container to load the new frames: docker compose restart"
echo "curl $HOSTNAME/$NAME"
exit 0
