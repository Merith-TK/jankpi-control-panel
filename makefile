default:
	go run ./
pi:
	GOOS=linux GOARCH=arm go build ./
	scp ./* pi@192.168.0.111:/home/pi/workspace/control-panel/
	