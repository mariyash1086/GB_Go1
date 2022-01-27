package HW7

//	1.	С какими интерфейсами мы уже сталкивались в предыдущих уроках?
//	Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.
//ОТВЕТ: например, fmt.scanLn, fmt.printLn,math.Pow,math.Sqrt

//	2.Посмотрите примеры кода в своём портфолио.
//	Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?
//ОТВЕТ: можно добавить дефер в калькуляторе при проверке деления на 0.
//defer func() {
//	errFromPanic := recover()
//	if errFromPanic != nil {
//		err = fmt.Errorf("can not divide by zero", errFromPanic)
//	}
//}

//func calculate(a int, b int){
// if b ==0 {
//	panic("panic in divide by zero")
//}
//}
