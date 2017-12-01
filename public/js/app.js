
$(document).ready(function(){
    var carts = [];
    $(".cart-btn").click(function(){
        var id = $(this).parent().find(".id").val();
        var name = $(this).parent().find(".name").val();
        var image = $(this).parent().find(".image").val();
        var price = $(this).parent().find(".price").val();
        item = [id, name, price,image];
        console.log(item[0]);
        carts.push(item);
        


            var cookieValue = $.cookie(id,carts);
            console.log( $.cookie("2"));
            $(".cart-infor-item").append('<li><div class="cart-img"><img src="/public/img/cart/1.png" alt=""></div><div class="cart-details"><a href="#">'+name+'</a><p>1 x $'+price+'</p></div><div class="btn-edit"></div><div class="btn-remove"></div></li>');

   
    });
});