<!-- checkout area start -->
<div class="checkout-area">
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <div class="location">
                    <ul>
                        <li><a href="index.html" title="go to homepage">Home<span>/</span></a>  </li>
                        <li><strong> checkout</strong></li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col-sm-3">
                <div class="product-sidebar">
                    <div class="sidebar-title">
                        <h2>Shopping Options</h2>
                    </div>
                    <div class="single-sidebar">
                        <div class="single-sidebar-title">
                            <h3>Category</h3>
                        </div>
                        <div class="single-sidebar-content">
                            <ul>
                                <li><a href="#">Dresses (4)</a></li>
                                <li><a href="#">shoes (6)</a></li>
                                <li><a href="#">Handbags (1)</a></li>
                                <li><a href="#">Clothing (3)</a></li>
                            </ul>
                        </div>
                    </div>
                    <div class="single-sidebar">
                        <div class="single-sidebar-title">
                            <h3>Color</h3>
                        </div>
                        <div class="single-sidebar-content">
                            <ul>
                                <li><a href="#">Black (2)</a></li>
                                <li><a href="#">Blue (2)</a></li>
                                <li><a href="#">Green (4)</a></li>
                                <li><a href="#">Grey (2)</a></li>
                                <li><a href="#">Red (2)</a></li>
                                <li><a href="#">White (2)</a></li>
                            </ul>
                        </div>
                    </div>
                    <div class="banner-left">
                        <a href="#">
                            <img src="/static/frontend/img/product/banner_left.jpg" alt="">
                        </a>
                    </div>
                </div>
            </div>
            <div class="col-sm-9">
                <div class="checkout-banner hidden-xs">
                    <a href="#">
                        <img src="/static/frontend/img/checkout/checkout_banner.jpg" alt="">
                    </a>
                </div>
                <div class="checkout-heading">
                    <h2>Checkout</h2>
                </div>
                <div class="checkout-accordion">
                    <div class="panel-group" id="accordion" role="tablist" aria-multiselectable="true">
                        <div class="panel panel-default">
                            <div class="panel-heading" role="tab" id="headingFive">
                                <h4 class="panel-title">
                                    <a class="collapsed" role="button" data-toggle="collapse" data-parent="#accordion" href="#collapseFive" aria-expanded="false" aria-controls="collapseFive">
                                        Checkout: Payment Method
                                        <i class="fa fa-caret-down"></i>
                                    </a>
                                </h4>
                            </div>
                            <form action="/checkout" method="POST">
                                
                                <div id="collapseFive" class="panel-collapse collapse" role="tabpanel" aria-labelledby="headingFive">
                                    <div class="panel-body">
                                        <div class="patment-method">
                                            <p>Please select the preferred payment method to use on this order.</p>
                                            <div class="radio">
                                                <input type="radio" checked="" value="visa">Visa
                                            </div>
                                            VAT: 5$ </br>
                                            <h2> Total: <span style="color:red;"> </span>$</h2>
                                            <input class="id" type="hidden" value="visa" name="menthod"> </input>
                                            <input class="id" type="hidden" value="" name="order_id"> </input>
                                            <input class="id" type="hidden" value="5" name="vat"> </input>
                                            <input class="id" type="hidden" value="" name="total"> </input>
                                            <p> <strong> Add Address For Your Order</strong></p>
                                            <p> <textarea name="address" rows="8"></textarea> </p>
                                            <div class="privacy-policy">
                                                I have read and agree to the
                                                <a href="#">Privacy Policy</a>
                                                <input checked type="checkbox" name="agree">
                                                <button type="submit" value="Continue" class="check-button">Pay</button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<!-- checkout -->