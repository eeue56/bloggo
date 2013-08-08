$(document).ready(function() {
    
    var createPost = function(data) {
        var tpl = "
        <h2 class='title'>{{Title}}</h2>
        <h5 class='date'>{{Time}}</h5>
        <div class='body'>
            {{Content}}
        </div>";

        return Mustache.render(tpl, data);
    };

    var _post = function(url, id){
        var data = {};
        $.ajax({
            type: "POST",
            async: false,
            url: url + id
        }).done(function(data) {
            createPost(data);
        });
    };

    var nextPost = function(id) {
        _post("NextPost/", id);
    };

    var previousPost = function(id) {
        _post("PreviousPost/", id);
    };
});