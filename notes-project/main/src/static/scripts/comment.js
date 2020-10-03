
function update(id){
    $.ajax("http://localhost:2020/api/updateNote",{
        type:"POST",
        data:{"id":id, "Text":$("#text").val(),
        },
        success: function(){
            window.alert("Saved!")
        }

    })
}