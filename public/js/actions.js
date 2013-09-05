$(document).ready(function() {
	live();
});

$(document).keydown(function(e){
	if(e.keyCode == 39) {
		nextPost();
	} else if(e.keyCode == 37) {
		lastPost();
	}
});

// current -> before
// before -> after [l]
// after -> current
var nextPost = function() {
	console.log("next")
	$(".posts").animate({
		"left": "+=100%"
	});
};

// current -> after
// before -> current
// after -> before [l]
var lastPost = function() {
	$(".posts").animate({
		"left": "-=100%"
	});
};

var live = function() {
	$(".post").on("click", lastPost);
};