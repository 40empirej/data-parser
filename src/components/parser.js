const fs = require('fs');
const csv = require('csv-parser');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  parse() {
    const data = [];
    const stream = fs.createReadStream(this.filePath)
      .pipe(csv());

    stream.on('data', (row) => {
      data.push(row);
    });

    return new Promise((resolve) => {
      stream.on('end', () => {
        resolve(data);
      });
    });
  }
}

module.exports = Parser;