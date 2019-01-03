[![Build Status](https://travis-ci.org/bhowmikp/sudoku-solver-server.svg?branch=master)](https://travis-ci.org/bhowmikp/sudoku-solver-server)

# [Sudoku Solver Server](https://dry-hollows-45767.herokuapp.com/)
Solves sudoku puzzles. This code is written for the server.

## Installation

```bash
./build-and-run.sh
```

## Usage

- go to localhost:8000/?board=[board]
- board consists of 81 spots. Only 1-9 are valid, everything else treated as blank spots

## Deployment

Deploy to Heroku:
```sh
heroku container:push web
heroku container:release web
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

Code inside container:
```sh
sudo docker run -it golang-app /bin/bash
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
