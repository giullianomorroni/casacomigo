var HOST = location.host;
var PROTOCOL = window.location.protocol;

$(document).ready(function() {

	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/produto/listar/limite/12/pular/0",
		  success: function(data) { mountProducts(data); },
		  dataType: "json"
	});

});


function mountProducts(data) {
	var div = $("#listView");
	for (var d in data) {	
		produto =data[d]
		var html = 
			"<div class='row'> "+
			"	<div class='span2'> <img src='/public/store/img/"+produto.Codigo+"/thumb.jpg' alt=''/> </div> "+
			"	<div class='span4'> "+
			"		<h3>Disponível</h3> "+
			"		<hr class='soft'/> "+
			"		<h5>" + produto.Titulo + " </h5> "+
			"		<p> " + produto.Descricao + " </p> "+
			"		<a class='btn btn-small pull-right' href='/loja/produto/"+produto.Codigo+"'>Mais Informações</a> "+
			"		<br class='clr'/> "+
			"	</div> "+
			"	<div class='span3 alignR'> "+
			"		<h3> "+produto.Preco+" </h3> "+
			"		<label class='checkbox'> "+
			"		</label><br/> "+
			"		<a class='btn btn-large btn-primary'> Adicionar ao Carrinho <i class='icon-shopping-cart'></i></a> "+
			"		<a href='/loja/produto/"+produto.Codigo+"' class='btn btn-large'><i class='icon-zoom-in'></i></a> "+
			"	</div> "+
			"</div> "+
			"<hr class='soft'/>";
		$(html).appendTo(div);

		var ul = $("#blockView_ul")
		var html = 
			"<li class='span3'> "+
			"	<div class='thumbnail'> "+
			"		<a href='/loja/produto/"+produto.Codigo+"'><img src='/public/store/img/"+produto.Codigo+"/thumb.jpg' alt=''/></a> "+
			"		<div class='caption'> "+
			"			<h5> </h5> "+
			"			<p></p> "+
			"			<h4 style='text-align:center'> "+
			"				<a class='btn' href='/loja/produto/"+produto.Codigo+"'> <i class='icon-zoom-in'></i></a>  "+
			"				<a class='btn' href='#'>Adicionar ao Carrinho <i class='icon-shopping-cart'></i></a>  "+
			"				<a class='btn btn-primary' href='#'> R$ " +produto.Preco+ "</a>"+
			"			</h4> "+
			"		</div> "+
			"	</div> "+
			"</li>";
		$(html).appendTo(ul);
	}
}
