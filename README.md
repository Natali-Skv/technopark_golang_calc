# technopark_golang_calc
индивидуальное задание 1.2, дисциплина "Разработка веб-сервисов на Golang"

Нужно написать калькулятор, умеющий вычислять выражение, подаваемое на STDIN.  
Тут также нужны тесты. Тестами нужно покрыть все операции и несколько функций.

---

<details>
  <summary>:red_circle: Развернуть список 11 доступных констант!</summary>

- **e**
- **pi**
- **phi**
- **sqrt2**
- **sqrte**
- **sqrtpi**
- **sqrtphi**
- **ln2**
- **log2e**
- **ln10**
- **log10e**
</details>

---

<details>
  <summary>:red_circle: Развернуть список 48 доступных функций!</summary>

- **abs** absolute value of x
- **acos** arccosine, in radians, of x
- **acosh** inverse hyperbolic cosine of x
- **asin** arcsine, in radians, of x
- **asinh** inverse hyperbolic sine of x
- **atan** arctangent, in radians, of x
- **atan2** arc tangent of y/x, using the signs of the two to determine the quadrant of the return value
- **atanh** inverse hyperbolic tangent of x
- **cbrt** cube root of x
- **ceil** least integer value greater than or equal to x
- **copysignalue** with the magnitude of x and the sign of y
- **cos** cosine of the radian argument x
- **cosh** hyperbolic cosine of x
- **dim** maximum of x-y or 0
- **erf** error function of x
- **erfc** complementary error function of x
- **erfcinv** inverse of Erfc(x)
- **erfinv** inverse error function of x
- **expx**, the base-e exponential of x
- **exp2x**, the base-2 exponential of x
- **expm1x** - **1,** the base-e exponential of x minus 1
- **floor** greatest integer value less than or equal to x
- **gamma** Gamma function of x
- **hypott**(p*p + q*q)
- **j0** order-zero Bessel function of the first kind
- **j1** order-one Bessel function of the first kind
- **log** natural logarithm of x
- **log10** decimal logarithm of x
- **log1p** natural logarithm of 1 plus its argument x
- **log2** binary logarithm of x
- **logb** binary exponent of x
- **max** larger of x or y
- **min** smaller of x or y
- **mod** floating-point remainder of x/y
- **nanIEEE** 754 “not-a-number” value
- **nextafter** next representable float64 value after x towards y
- **powy**, the base-x exponential of y
- **remainder** IEEE 754 floating-point remainder of x/y
- **round** nearest integer, rounding half away from zero
- **roundtoeven** nearest integer, rounding ties to even
- **sin** sine of the radian argument x
- **sinh** hyperbolic sine of x
- **sqrt** square root of x
- **tan** tangent of the radian argument x
- **tanh** hyperbolic tangent of x
- **trunc** integer value of x
- **y0** order-zero Bessel function of the second kind
- **y1** order-one Bessel function of the second kind
</details>

---


### Пример работы
![alt-text](https://github.com/Natali-Skv/technopark_golang_calc/blob/dev-with-fcns/example.gif)

```bash

calculator>> (2+2)*2
8
calculator>> (2 + 2) * 2
8
calculator>> log10(1000) + sqrt2
4.414213562373095
calculator>> abs( -pow(2,3) + sin(pi) - log2(1024))
18
calculator>> log(pow(e,10))
10
calculator>> cbrt(700 + 29) -1.1
7.9
calculator>> max(-trunc(e), -trunc(pi))
-2
calculator>> sin(asin(sqrt2/2))
0.7071067811865475
calculator>> (1+2
expression is not valid
calculator>> 

```
