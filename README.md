# fortunes
A collection of fortune database files for me

## 安装 & 配置

1. 安装 fortune 
	
	```
	brew install fortune
	```

2. 下载安装数据文件

	```
	git clone git@github.com:kliu/fortunes.git
	
	cp songci-fortunes* /usr/local/share/games/fortunes/
	```

3. 在 ~/.bashrc 或 ~/.zshrc 文件的结尾，加上下面几行。

	```
	echo
	echo "=============== Poem Of The Day ==============="
	echo
	fortune -e songci-fortunes
	echo
	echo "================================================"
	echo
	```

## 生成数据文件

1. 所有条目都写入一个文本文件，文件名任意，不加扩展名。

2. 条目之间用单独一行的百分号（%）分隔。

	```
	疏綺籠寒，淺雲棲月。
	- 丁宥 · 水龍吟
	%
	雁風吹裂雲痕，小樓一線斜陽影。
	- 丁宥 · 法曲獻仙音
	%
	蟬碧句花，雁紅攢月。
	- 丁察院 · 萬年歡
	%
	```

3. 创建索引文件

	```
	strfile <file> <file.dat>
	```

## 快捷方式

如果安装了 Golang 的开发环境，运行下面命令即可完成以上说的全部。

```
go run songci.go
```
