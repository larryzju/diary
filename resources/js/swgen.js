function generate_navigator_ulist(lists) {
	var ul = $('<ul></ul>');
	for (var ele of lists) {
		if ( ele.children ) {
			summary = $('<summary>').append(ele.title);
			details = $('<details>').append(summary, generate_navigator_ulist(ele.children));
			ul.append($('<li>').append(details));
		} else {
			anchor = $('<a>').attr('href', ele.link).append(ele.title);
			ul.append($('<li>').append(anchor));
		}
	}
	return ul
}

$(document).ready(function() {
	const root = '/diary';
	$.getJSON(root + '/nav.json', function(nav) {
		ul = generate_navigator_ulist(nav.children);
		$('nav#content').append(ul);

		// highlight the link
		let urlpath = decodeURI(location.pathname);

		// show current links
		let link = $(`nav#content li a[href="${urlpath}"]`)
		link.addClass("li-focus");
		link.parents("details").attr("open",true);
		link[0].scrollIntoView();
	})
})
