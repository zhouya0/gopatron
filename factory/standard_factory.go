package factory

// 这个factory要达到这样一种效果：客户端实例化一个基类，然后根据不同的
// 业务要求来实例化不同的子类。

// 不难看出，这个接口里面，是Result函数来实现不同的业务逻辑。
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 这个就是factory的核心代码，只要实现了Create函数返回一个Opetator就可以了。
type OperatorFactory interface {
	Create() Operator
}

// 写好interface之后，我们应该马上想到，interface的不同实现必须依赖于不同的struct，
// 所以在这里，至少有两个Operator的struct(当然实现的角度来看，其实是一个OperatorBase，然后两个struct继承它)，

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 这里开始继承这个基类

// 加法
type PlusOperator struct {
	*OperatorBase
}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

// 减法
type MinusOperator struct {
	*OperatorBase
}

func (m MinusOperator) Result() int {
	return m.a - m.b
}

// 然后创建不同的struct来实现不同的factory方法

// 加法
type CreatePlus struct{}

// 这里有个坑，那就是go里面的继承像是继承又不像，因为在实例化子类的时候，还是需要实例化父类。
func (c CreatePlus) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 减法
type CreateMinus struct{}

func (c CreateMinus) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
