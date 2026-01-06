const fs = require('fs');
const path = require('path');
const xml2js = require('xml2js');

class Parser {
  constructor() {
    this.parser = new xml2js.Parser();
  }

  parseFile(filePath) {
    return new Promise((resolve, reject) => {
      fs.readFile(filePath, (err, data) => {
        if (err) {
          reject(err);
        } else {
          this.parser.parseString(data, (err, result) => {
            if (err) {
              reject(err);
            } else {
              resolve(result);
            }
          });
        }
      });
    });
  }
}

module.exports = Parser;