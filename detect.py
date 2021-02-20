#!/usr/bin/python3

from glob import glob
import numpy as np
import os
import PIL
import PIL.Image
import tensorflow as tf
import tensorflow_datasets as tfds
from tensorflow.keras import layers
from tensorflow import keras

model = tf.keras.models.load_model('modely/my_model')
print(model.summary())
pictures = glob("data/*")
pictures.sort()
out = ""
i = 0

for p in pictures:
    if i % 16 == 0:
        out += f"{int(i/16):02}.pbm\n"
    img = keras.preprocessing.image.load_img(
        p, target_size=(120,120)
    )
    img_array = keras.preprocessing.image.img_to_array(img)
    img_array = tf.expand_dims(img_array, 0) # Create a batch
    predictions = model.predict(img_array)
    score = tf.nn.softmax(predictions[0])
    if np.argmax(score) == 0:
        out += ":("
    else:
        out += ":)"
    if (i+1)%4 == 0:
        out += "\n"
    else:
        out += " "
    i += 1
with open("output.txt", mode="w", encoding="utf-8") as f:
    f.write(out)
