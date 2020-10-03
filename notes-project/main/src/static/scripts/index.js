function del(id) {
    $.ajax("http://localhost:2020/api/delete",{
        type:"POST",
        data:{"id":id},
    })
}