$(document).ready(function() {
    $(".submit").click(function() {
        console.log("Here");
        var data = {
            "Content" : $(".body").text().trim(),
            "Title" : $("#title").text().trim()
        };
        $.post("Create", JSON.stringify(data));
    });
});
