bowling-go-kata
===============

• Cada juego incluye 10 turnos para el jugador.

• En cada turno, el jugador tiene hasta 2 intentos para tirar todos los pinos.

• El puntaje del juego es la suma de los puntajes de todos los turnos.

• Si en los 2 intentos, el jugador… – …no logra tirar todos los pinos, su puntaje para ese turno será la suma de los pinos tirados. – …logra tirar todos los pinos, esto se llama un “Spare” y su puntaje para ese turno será de 10 + la cantidad de pinos tirados en su siguiente intento.

• Si en el primer intento de su turno tira todos los pinos, se llama un “Strike”. – Esto termina el turno del jugador. – Y su puntaje para ese turno será de 10 + el total simple de pinos tirados en sus próximos 2 intentos
Bowling Kata

• Si obtiene un “Spare” o “Strike” en el último (décimo) turno, el jugador puede tirar 1 o 2 intentos más como bono, respectivamente. – Estos intentos del bono se consideran parte del mismo turno. – Si en alguno de los intentos del bono hace un “Strike”, el proceso no se repite. – Los intentos del bono solo se utilizan para calcular el puntaje total del turno final.

Casos de prueba sugeridos

• X es un strike, / es un spare y - es nada.

• 12 intentos: 12 strikes. XXXXXXXXXXXX = 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 + 10+10+10 = 300

• 20 intentos: 10 pares de 9 y nada. 9-9-9-9-9-9-9-9-9-9- = 9 + 9 + 9 + 9 + 9 + 9 + 9 + 9 + 9 + 9 = 90

• 21 intentos: 10 pares de 5 y spare + 5. 5/5/5/5/5/5/5/5/5/5/5 = 10+5 + 10+5 + 10+5 + 10+5 + 10+5 + 10+5 + 10+5 + 10+5 + 10+5 + 10+5 = 150

Referencias

http://www.solveet.com/exercises/Kata-Bowling/10/

http://www.youtube.com/watch?v=laSVXHFEJXQ

http://www.youtube.com/watch?v=q4-4x6jWJYU

http://www.youtube.com/watch?v=6tPK9SYXdPY
