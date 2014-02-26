var page = require('./libs/page.js');
//var ractive = require('./views/test.js');



console.log('app.js');


page('*', function(ctx, next) {
//	ractive.set({title: 'win'});
	console.log('path * :',ctx.path);
	console.log('hash:',ctx.hash);
	console.log('ctx',ctx);
	console.log('app js page nav * function');
});

var section = function section(ctx,next) {
	console.log('section');
	console.log('path:',ctx.path);
	console.log('hash:',ctx.hash);
}


page();