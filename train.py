#!/usr/bin/python3

import numpy as np
import os
import PIL
import PIL.Image
import tensorflow as tf
import tensorflow_datasets as tfds
from tensorflow.keras import layers
from tensorflow import keras

num_classes = 2
AUTOTUNE = tf.data.AUTOTUNE

train_ds = tf.keras.preprocessing.image_dataset_from_directory("model", validation_split=0.2, subset="training", seed=123, image_size=(120, 120), label_mode='int')
val_ds = tf.keras.preprocessing.image_dataset_from_directory("model", validation_split=0.2, subset="validation", seed=123, image_size=(120, 120), label_mode='int')

train_ds = train_ds.cache().prefetch(buffer_size=AUTOTUNE)
val_ds = val_ds.cache().prefetch(buffer_size=AUTOTUNE)


#model = tf.keras.Sequential([
#  layers.experimental.preprocessing.Rescaling(1./255),
#  layers.Conv2D(32, 3, activation='relu'),
#  layers.MaxPooling2D(),
#  layers.Conv2D(32, 3, activation='relu'),
#  layers.MaxPooling2D(),
#  layers.Conv2D(32, 3, activation='relu'),
#  layers.MaxPooling2D(),
#  layers.Flatten(),
#  layers.Dense(128, activation='relu'),
#  layers.Dense(num_classes)
#])

model = keras.Sequential(
	[
	layers.experimental.preprocessing.RandomFlip("horizontal", input_shape=(120, 120, 3)),
	layers.experimental.preprocessing.RandomRotation(0.1),
	layers.experimental.preprocessing.Rescaling(1./255),
	layers.Conv2D(32, 3, activation='relu'),
	layers.MaxPooling2D(),
	layers.Conv2D(32, 3, activation='relu'),
	layers.MaxPooling2D(),
	layers.Conv2D(32, 3, activation='relu'),
	layers.MaxPooling2D(),
	layers.Flatten(),
	layers.Dense(128, activation='relu'),
	layers.Dense(num_classes),
	]
)

model.compile(
  optimizer='adam',
  loss=tf.losses.SparseCategoricalCrossentropy(from_logits=True),
  metrics=['accuracy'])

model.fit(
  train_ds,
  validation_data=val_ds,
  epochs=10
)
model.save('modely/my_model')
