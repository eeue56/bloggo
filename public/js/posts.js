$(document).ready(function() {

    var _post = function(text, id){
        var data = {};
        $.post(text + id, function(stuff)){
            data = stuff;
        };
        return data;
    };

    var nextPost = function(id) {
        return _post("NextPost/", id);
    };

    var previousPost = function(id) {
        return _post("PreviousPost/", id);
    };
});