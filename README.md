# Notes about golang

#### imports

- factored import statementを利用しよう
  - ()でグループでまとめる感じ
- [A Tour of GoのImports](https://go-tour-jp.appspot.com/basics/2)

```
# bad
import "fmt"
import "math"

# good
import (
	"fmt"
	"math"
)
```

#### "naked" return

- メソッドの中が長いときは、"naked" returnをやめる
- 長さの感覚は人によるかもしれない
- 20行を基準として考える記事が多かったような気がする(リーダブルコードもそうだったかは忘れた。。)
- Rails周りで基本3行みたいな記事もどこかで見た気がする
- Goだと。。。難しそう

```
# bad
func lognfunc(sum int) (x, y int) {
...長い
	return
}

# good
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

#### 初期化子(initializer)

```
# same
	var i, j int = 1, 2
	i, j := 1, 2

# but can't
k := 5
func main() {
...    
}
```

#### 特別な理由がない限りint32などは避けてintを使う

- 32-bitシステムではintは32bit
- 64-bitシステムではintは64bit
- [A Tour of GoのBasic types](https://go-tour-jp.appspot.com/basics/11)

#### 変数に初期値を与えずに宣言するとzero valueが与えられる

- int,floatは0
- boolはfalse
- stringは空文字''

#### 型変換は明示的に

```
# same
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

i := 42
f := float64(i)
u := uint(f)
```

#### しかし、型推論される

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

#### 定数はconstで宣言、 `:=` は無理

#### 64-bitでは表現が難しい時、定数使おすすめ
- [Tour of GoのNumeric Constants](https://go-tour-jp.appspot.com/basics/16)

#### ifステートメントで宣言した変数はifのスコープ内のみ

- else句でも使えるb

```
	if v := math.Pow(x, n); v < lim {
		return v
	}
```

#### deferで呼び出し元の最後に評価

- じゃあいつ書くの？
- 今でしょ（冗談）
- パッと思いつくのはfile.close()系
- deferを複数使った時、LIFOでrunされる

#### 配列の初期化

- 固定長

```
	primes := [6]int{2, 3, 5, 7, 11, 13}
```

#### スライスが可変長

- 配列よりスライスのほうが一般的らしい
- 配列とスライスの見分け方
  - `[n]T 型` or `[]T 型`
  
#### スライスの初期化

```
	[]bool{true, true, false}

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
```

#### 配列からスライスを作成時、上限/下限は省略できる

```
var a [10]int

# same
a[0:10]
a[:10]
a[0:]
a[:]
```

#### 組み込みのmake関数でスライスを作成

- capは容量
- lenは長さ

```
	a := make([]int, 5)
#=> a len=5 cap=5 [0 0 0 0 0]
```
