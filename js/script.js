// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates a random number generator

"use strict"

let randomNumber = Math.floor(Math.random() * 6) + 1

function myButtonClicked() {
  const numberGuessed = parseInt(
    document.getElementById("guessed-number").value
  )

  if (numberGuessed == randomNumber) {
    document.getElementById("answer").innerHTML = "Good job!"
    randomNumber = Math.floor(Math.random() * 6) + 1
  } else {
    document.getElementById("answer").innerHTML = "Try again."
  }
}


