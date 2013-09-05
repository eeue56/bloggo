$(document).ready(function() {
    $(".submit").click(function() {
        console.log("Here");
        var data = {
            "Content" : $(".body").html().trim(),
            "Title" : $("#title").html().trim()
        };
        console.log(data);
        var str = JSON.stringify(data)
        console.log(str);


        $.ajax({
            url : "Create", 
            data: str,
            dataType: "json",
            type: "post"
        });
    });
});
