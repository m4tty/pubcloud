


var Test = function() {

	console.log('test model constructor function');
	var self = this;

	this.title = 'Loading...';

	jQuery.ajax({
		type: 'GET',
		url: 'http://localhost:8080/pubs',
		success: function(data) {
			console.log(data);
			console.log('asdfasdf');
			self.title = decodeURIComponent(data);
		},
		//dataType: 'jsonp',
		crossDomain: false
	});
};




module.exports = Test;
