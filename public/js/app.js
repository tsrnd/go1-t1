
$(document).ready(function(){

    $(".cart-btn").click(function(event){
        event.preventDefault();
        var id = $(this).parent().find(".id").val();
        var name = $(this).parent().find(".name").val();
        var image = $(this).parent().find(".image").val();
        var price = $(this).parent().find(".price").val();
        var quantity = $(this).parent().find(".quantity").val();
        shoppingCart.addItemToCart(id,name, price, 1,image);
        displayCart();
    });
    function displayCart() {
        var cartArray = shoppingCart.listCart();
        console.log(cartArray);
            var output = "";
        for (var i in cartArray) {
            output += '<li><div class="cart-img"><img src="/public/img/product/'+cartArray[i].image+'" alt=""></div><div class="cart-details"><a href="#">'+cartArray[i].name+'</a><p>'+cartArray[i].count+' x $'+cartArray[i].price+'</p></div><div class="btn-edit"></div><div class="btn-remove"></div></li>';
        }
        $(".cart-infor-item").html(output);
        $(".sum-item").text(shoppingCart.countCart());
        $(".total").html( shoppingCart.totalCart() );
    }
    displayCart();
});
