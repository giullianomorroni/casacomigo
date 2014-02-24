var HOST = location.host;
var PROTOCOL = window.location.protocol;

$(document).ready(function() {
	
	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/produto/listar/limite/4/pular/0",
		  success: function(data) { mountFeaturedProducts(data, "#first_featured_products_ul"); },
		  dataType: "json"
	});

	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/produto/listar/limite/4/pular/4",
		  success: function(data) { mountFeaturedProducts(data, "#second_featured_products_ul"); },
		  dataType: "json"
	});

	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/produto/listar/limite/4/pular/8",
		  success: function(data) { mountFeaturedProducts(data, "#third_featured_products_ul"); },
		  dataType: "json"
	});

	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/produto/listar/limite/6/pular/0",
		  success: function(data) { mountFeaturedProducts(data, "#latest_products_ul"); },
		  dataType: "json"
	});

});

function mountFeaturedProducts(data, ul_id) {
	var ul = $(ul_id);
	for (var d in data) {	
		produto =data[d]
		var html = 
			"<li class='span3'> "+
			"<div class='thumbnail'> "+
			"	<i class='tag'></i> "+
			"	<a href='/loja/produto/"+produto.Codigo+"'><img src='/public/store/img/"+produto.Codigo+"/thumb.jpg' alt=''></a> "+
			"	<div class='caption'> "+
			"		<h5> "+produto.Titulo+"</h5> "+
			"	  	<h4><a class='btn' href='/loja/produto/"+produto.Codigo+"'> Detalhes </a> <span class='pull-right'> R$"+produto.Preco+" </span></h4> "+
			"	</div> "+
			"</div> "+
			"</li>";
		$(html).appendTo(ul);
	}
}

function mountLatestProducts(data, ul_id) {
	var ul = $(ul_id);
	for (var d in data) {	
		produto=data[d]
		var html = 
		"<li class='span3'>"+
		"	<div class='thumbnail'>"+
		"		<a  href='/loja/produto/"+produto.Codigo+"'><img src='/public/store/img/"+produto.Codigo+"/thumb.jpg' alt=''/></a>"+
		"			<div class='caption'>"+
		"				<h5>"+produto.Titulo+"</h5>"+
		"				<p> "+produto.Descricao.subString(0, 20)+"... </p>"+
		"				<h4 style='text-align:center'>" +
		"					<a class='btn' href='/loja/produto/"+produto.Codigo+"'> <i class='icon-zoom-in'></i></a> " +
		"					<a class='btn' href='#'>Adicionar ao Carrinho<i class='icon-shopping-cart'></i></a> " +
		"					<a class='btn btn-primary' href='#'>R$"+produto.Preco+"</a></h4>"+
		"			</div>"+
		"		</div>"+
		"</li>";
		$(html).appendTo(ul);
	}
}
