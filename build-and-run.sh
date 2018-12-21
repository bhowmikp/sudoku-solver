sudo docker build -t golang-app .
sudo docker run -it -p 8000:8000 --rm --name sudoku-app golang-app
