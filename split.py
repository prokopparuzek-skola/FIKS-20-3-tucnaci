#!/usr/bin/python3

from glob import glob
from PIL import Image

pictures = glob("input/*")
pictures.sort()
i = 0

for p in pictures:
    img = Image.open(p)
    for y in range(0, 480, 120):
        for x in range(0, 480, 120):
            cropImg = img.crop((x, y, x+120, y+120))
            cropImg.save(f"data/{i:04}.pbm")
            i += 1
