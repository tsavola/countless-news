// Suggest additional news sources by sending email to sources@countless.news.
//
// Criteria:
//
// - Focus on national news
// - Local news agencies preferred over global ones
// - In English
// - Atom or RSS feed

package main

var NationalSources = []Nation{
	{Name: "Austria", Flag: "🇦🇹", Sources: []Source{
		S{"https://www.thelocal.at/feeds/rss.php"},
	}},

	{Name: "Belgium", Flag: "🇧🇪", Sources: []Source{
		S{"https://deredactie.be/cm/vrtnieuws.english?mode=atom"},
	}},

	{Name: "Bulgaria", Flag: "🇧🇬", Sources: []Source{
		S{"https://feeds.feedburner.com/TheSofiaGlobe"},
	}},

	{Name: "Croatia", Flag: "🇭🇷", Sources: []Source{
		S{"http://www.thedubrovniktimes.com/news/croatia?format=feed&type=atom"},
	}},

	{Name: "Cyprus", Flag: "🇨🇾", Sources: []Source{
		S{"http://cyprus-mail.com/feed/"},
	}},

	{Name: "Czech Republic", Flag: "🇨🇿", Sources: []Source{
		S{"http://www.radio.cz/feeds/rss/en/en.xml"},
	}},

	{Name: "Denmark", Flag: "🇩🇰", Sources: []Source{
		S{"https://www.thelocal.dk/feeds/rss.php"},
	}},

	{Name: "Estonia", Flag: "🇪🇪", Sources: []Source{
		SSubstring{S{"https://feeds.feedburner.com/TheBalticTimesNews"}, "estonia"},
	}},

	{Name: "Finland", Flag: "🇫🇮", Sources: []Source{
		S{"https://feeds.yle.fi/uutiset/v1/recent.rss?publisherIds=YLE_NEWS"},
	}},

	{Name: "France", Flag: "🇫🇷", Sources: []Source{
		S{"http://www.france24.com/en/france/rss"},
		S{"https://www.thelocal.fr/feeds/rss.php"},
	}},

	{Name: "Germany", Flag: "🇩🇪", Sources: []Source{
		S{"https://www.thelocal.de/feeds/rss.php"},
	}},

	{Name: "Greece", Flag: "🇬🇷", Sources: []Source{
		S{"http://greece.greekreporter.com/feed/"},
	}},

	{Name: "Hungary", Flag: "🇭🇺", Sources: []Source{
		S{"https://dailynewshungary.com/feed/"},
	}},

	{Name: "Iceland", Flag: "🇮🇸", Sources: []Source{
		S{"http://icelandmonitor.mbl.is/rss/"},
	}},

	{Name: "Ireland", Flag: "🇮🇪", Sources: []Source{
		S{"http://feeds.breakingnews.ie/bnireland"},
	}},

	{Name: "Italy", Flag: "🇮🇹", Sources: []Source{
		S{"https://www.thelocal.it/feeds/rss.php"},
	}},

	{Name: "Latvia", Flag: "🇱🇻", Sources: []Source{
		SSubstring{S{"https://feeds.feedburner.com/TheBalticTimesNews"}, "latvia"},
	}},

	{Name: "Lithuania", Flag: "🇱🇹", Sources: []Source{
		SSubstring{S{"https://feeds.feedburner.com/TheBalticTimesNews"}, "lithuania"},
	}},

	{Name: "Luxembourg", Flag: "🇱🇺", Sources: []Source{
		S{"https://www.wort.lu/rss/en"},
	}},

	{Name: "Malta", Flag: "🇲🇹", Sources: []Source{
		S{"http://www.maltatoday.com.mt/rss/news/"},
	}},

	{Name: "Netherlands", Flag: "🇳🇱", Sources: []Source{
		S{"https://www.dutchnews.nl/feed/?news"},
	}},

	{Name: "Norway", Flag: "🇳🇴", Sources: []Source{
		S{"https://www.thelocal.no/feeds/rss.php"},
	}},

	{Name: "Poland", Flag: "🇵🇱", Sources: []Source{
		S{"http://www.thenews.pl/Rss/47e92f6a-fbb8-4bb3-9b2e-128125c4b722"},
	}},

	{Name: "Portugal", Flag: "🇵🇹", Sources: []Source{
		S{"https://portugalresident.com/articles.xml"},
	}},

	{Name: "Romania", Flag: "🇷🇴", Sources: []Source{
		S{"https://www.romania-insider.com/feed/atom/"},
	}},

	{Name: "Russia", Flag: "🇷🇺", Sources: []Source{
		S{"https://themoscowtimes.com/feeds/news.xml"},
	}},

	{Name: "Slovakia", Flag: "🇸🇰", Sources: []Source{
		S{"https://spectator.sme.sk/rss-title"},
	}},

	{Name: "Slovenia", Flag: "🇸🇮", Sources: []Source{
		S{"http://www.sloveniatimes.com/rss?category_id=1"},
	}},

	{Name: "Spain", Flag: "🇪🇸", Sources: []Source{
		S{"https://www.thelocal.es/feeds/rss.php"},
	}},

	{Name: "Sweden", Flag: "🇸🇪", Sources: []Source{
		S{"https://www.thelocal.se/feeds/rss.php"},
	}},

	{Name: "Switzerland", Flag: "🇨🇭", Sources: []Source{
		S{"https://www.thelocal.ch/feeds/rss.php"},
	}},

	{Name: "United Kingdom", Flag: "🇬🇧", Sources: []Source{
		S{"https://feeds.bbci.co.uk/news/uk/rss.xml"},
		S{"https://www.independent.co.uk/news/uk/rss"},
		S{"https://www.theguardian.com/uk-news/rss"},
	}},
}
