
$(document).ready(function(){

    $(".cart-btn").click(function(event){
        event.preventDefault();
        var id = $(this).parent().find(".id").val();
        var name = $(this).parent().find(".name").val();
        var image = $(this).parent().find(".image").val();
        var price = $(this).parent().find(".price").val();
        var quantity = parseInt($(this).parent().find(".quantity").val(),10);
        shoppingCart.addItemToCart(id,name, price, quantity,image);
        displayCart();
    });
    function displayCart() {
        var cartArray = shoppingCart.listCart();
        console.log(cartArray);
            var output = "";
        for (var i in cartArray) {
            output += '<li><div class="cart-img"><img style="width:50px; height:50px;" src="/public/img/product/'+cartArray[i].image+'" alt=""></div><div class="cart-details"><a href="/single-product/'+cartArray[i].id+'">'+cartArray[i].name+'</a><p>'+cartArray[i].count+' x $'+cartArray[i].price+'</p></div><div class="btn-edit"></div><div data-name="'+cartArray[i].name+'" class="btn-remove"></div></li>';
        }
        $(".cart-infor-item").html(output);
        $(".sum-item").text(shoppingCart.countCart());
        $(".total").html( shoppingCart.totalCart() );
    }

    function showTable(){
        var cartArray = shoppingCart.listCart();
        
            var item = "";
        for (var i in cartArray) {
            item += '<tr class="item-table"><td class="cart-item-img"><a href="single-product.html"><img src="/public/img/cart/'+cartArray[i].image+'" alt=""></a></td><input type="hidden" name="productID[i]" value="'+cartArray[i].id+'"><td class="cart-product-name"><a href="single-product.html">'+cartArray[i].name+'</a></td><td class="unit-price"><span>'+cartArray[i].price+'</span><input type="hidden" value="'+cartArray[i].price+'" name="unit-price[i]"></td><td class="quantity"><span>'+cartArray[i].count+'</span><input type="hidden" value="'+cartArray[i].count+'" name="quantity[i]"></td><td class="subtotal"><span>'+cartArray[i].price*cartArray[i].count+'</span><input type="hidden" value="'+cartArray[i].total*cartArray[i].image+'" name="subtotal[i]"></td><td class="remove-icon"><input type="hidden" value="'+cartArray[i].name+'" class="pro-name"><a href="#"><img class="table-remove" data-name="'+cartArray[i].name+'" src="/public/img/cart/btn_remove.png" alt=""></a></td></tr>';
        }
        $("#order-table").html(item);
        $(".order_subtotoal").html( shoppingCart.totalCart() );
    }
    displayCart();
    if (($("#order-table").length > 0)){
        showTable();
    }

    $(".table-remove").click(function(){
        var name = $(this).attr("data-name");
        shoppingCart.removeItemFromCartAll(name);
        showTable();
    });

    $(".btn-remove").click(function(){
        var name = $(this).attr("data-name");
        shoppingCart.removeItemFromCartAll(name);
        displayCart();
    });
    $(".check-button").click(function(){
        shoppingCart.clearCart();
        displayCart();
    });
});
