# dsBackuper
Простенький бекапер для сейвов игр серии Dark Souls. 
Нужные директории редактируются путём изменения массива sources(не забываем про бэкслеш в начале).
Для сборки .exe файла требуется, собсна, сам голанг с $GOPATH. Просто переходим в папку и выполняем go build -o filename.exe sourcename.go , где filename - имя выходного .exe файла, а sourcename.go - имя исходника.
Пример использования: go build -o dsBackup.exe main.go
