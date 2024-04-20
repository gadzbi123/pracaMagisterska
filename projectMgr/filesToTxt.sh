#!/bin/bash
base_dir="/mnt/g/baza_mgr/jacek/"
out_dir="/mnt/g/baza_mgr/sorted/"

programs_required=("rsync" "catdoc" "djvused" "pdftotext")
for program in ${programs_required[@]}; do
if ! [ -x "$(command -v $program)" ]; then
  echo Error: $program is not installed. >&2
  exit 1
fi
done

for file in $base_dir*; do
    if [ -f "$file" ]; then
        filename="${file%%.*}"
        ext="${file##*.}"
        ext_double="${file#*.}"
        if [ "$ext_double" == "tar.gz" ]; then
        echo run scipt again to decompress $base_dir$file
        fi
        case $ext in
        ".txt") rsync -ah --progress  $base_dir$file ${out_dir}txts/$filename.txt ;;
        ".pdf") pdftotext $base_dir$file ${out_dir}pdfs/$filename.txt ;;
        ".doc") catdoc $base_dir$file | tr -d '\n' > ${out_dir}docs/$filename.txt  ;;
        (".mp3" | ".MP3")) echo mp3 format is not supported;;
        ".djvu") djvused $base_dir$file -e print-pure-txt > ${out_dir}djvus/$filename.txt ;;
        *) echo "Unsupported format $ext in $base_dir$file";; 
        esac
    fi
done
