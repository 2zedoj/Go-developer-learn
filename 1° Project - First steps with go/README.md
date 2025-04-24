# Challenge: Thermometric Scale Conversion

  
Write a code in Go using all the knowledge acquired so far! And in this code you will need, based on Types of variables, objects, classes, interfaces and exceptions so that the program can solve the problem presented. In addition, we increase the level of complexity, which will include another type of thermometer scale and will give the user the opportunity to define what the value is, what its current scale is and what it wants to convert to.

## General Objective
Create a program in Go language that converts thermometric scales

## Problem Description
A high school teacher asked his students to develop a program to convert the boiling point of water from Kelvin to degrees Celsius.

## Learnings

To solve this challenge, I used the resources taught in previous classes, such as concepts of variable types, functions, objects, packages, constant, interfaces and exceptions.

Below, I will list the concepts learned and how they were incorporated into the project to solve this problem.

**1° Step** : First we create the scale package files, namely scale, kelvin, fahrenheit, celsius. In the scale file, there will be the object, which will help us to centralize the data of the current value, type of the current scale and the desired one, in addition to having the interface of the functions that the scales will have.

    - scale
	    - celsius.go
	    - fahrenheit.go
	    - kelvin.go
	    - scale.go


**2° Step** : Now with the files, in scale we have

 - Constant, of which there are types existing in our scale system
 - Structure, which centralizes the data that we will need from the user to perform the conversion, being the value, the current scale and the desired scale
 - Interface, which defines the functions that the scales will have to perform
 - GetScale function, which allows, with the scale type, to direct the scale functions
 - String() function, It only serves to show the name of the scale type
  

**3° Step** : In the scale files (Celsius, Fahrenheit and Kelvin), we create a structure, which allows us to direct the scale and we do not need to add any parameters, just to direct it correctly. Then we call the interface functions with the appropriate conversion, being the current scale and the scale that will be converted, for example the Fahrenheit scale, from which to convert its value to Celsius, you have to do the calculation below

    package  scale
    type  Fahrenheit  struct{}
    func(f Fahrenheit) ConvertCelsius(value  *float64) float64 {
	    return (*value-32)*5/9
	}

  

**4° Step** : In main, with the help of the bufio, os and string packages, we can get the information we need, such as value, current scale and the desired one. Always validating the data returned by the user.

**5° Step**: Finally, still in the main, we take the current scale and convert it to the desired scale.