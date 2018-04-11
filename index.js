var lineReader = require('readline').createInterface({
	input: require('fs').createReadStream('data3.txt')
  });
  
  var regex2 = new RegExp('\\w+');

  lineReader.on('line', function (line) {
	  	let dateStr = (new Date(parseInt(line))).toISOString();
		// console.log(dateStr)
	//   let dateStr = "2018-01-24T23:59:59.946Z";
	  if(dateStr.match(/.*T23:59.*/g)){
		console.log(">>>>>>>>>>>>> Found it");
		console.log(dateStr);
	  }
  });