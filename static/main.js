$(document).ready(function (e) {
	$(".sidebar").css("height", $(document).height())

	$("#search-bar").keyup(function (up) {
		$.post("/search/", { search: $("#search-bar").val() }, function (data) {
			$("div.content").html(data)
		})
	})


	$("#menu-button").click(function (click) {
		click.preventDefault()
		$(".sidebar").toggleClass("hide")
		$(".content").toggleClass("wide")
	})




})

