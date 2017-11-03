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

	{Name: "Albania", Flag: "🇦🇱", Sources: []Source{
		S{"http://www.tiranatimes.com/?feed=rss2"},
	}},

	{Name: "Andorra", Flag: "🇦🇩", Sources: []Source{
		S{"https://all-andorra.com/feed/"},
	}},

	{Name: "Armenia", Flag: "🇦🇲", Sources: []Source{
		S{"https://news.am/eng/rss/"},
	}},

	{Name: "Australia", Flag: "🇦🇺", Sources: []Source{
		S{"http://www.news.com.au/feed/"},
	}},

	{Name: "Austria", Flag: "🇦🇹", Sources: []Source{
		S{"https://www.thelocal.at/feeds/rss.php"},
	}},

	{Name: "Azerbaijan", Flag: "🇦🇿", Sources: []Source{
		S{"http://news.az/articles.rss"},
	}},

	{Name: "Belarus", Flag: "🇧🇾", Sources: []Source{
		S{"http://eng.belta.by/rss"},
	}},

	{Name: "Belgium", Flag: "🇧🇪", Sources: []Source{
		S{"https://deredactie.be/cm/vrtnieuws.english?mode=atom"},
	}},

	{Name: "Bosnia and Herzegovina", Flag: "🇧🇦", Sources: []Source{
		S{"http://www.telegraph.co.uk/news/rss.xml"},
		S{"https://www.nytimes.com/svc/collections/v1/publish/http://www.nytimes.com/topic/destination/bosnia-and-herzegovina/rss.xml"},
		S{"https://www.theguardian.com/world/bosnia-and-herzegovina/rss"},
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

	{Name: "East Timor", Flag: "🇹🇱", Sources: []Source{
		S{"https://www.theguardian.com/world/timor-leste/rss"},
	}},

	{Name: "Estonia", Flag: "🇪🇪", Sources: []Source{
		S{"https://www.theguardian.com/world/estonia/rss"},
	}},

	{Name: "Finland", Flag: "🇫🇮", Sources: []Source{
		S{"https://feeds.yle.fi/uutiset/v1/recent.rss?publisherIds=YLE_NEWS"},
	}},

	{Name: "France", Flag: "🇫🇷", Sources: []Source{
		S{"http://www.france24.com/en/france/rss"},
		S{"https://www.thelocal.fr/feeds/rss.php"},
	}},

	{Name: "Georgia", Flag: "🇬🇪", Sources: []Source{
		S{"http://www.civil.ge/eng/rss.php"},
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

	{Name: "Indonesia", Flag: "🇮🇩", Sources: []Source{
		S{"http://www.abc.net.au/news/feed/26574/rss.xml"},
	}},

	{Name: "Ireland", Flag: "🇮🇪", Sources: []Source{
		S{"http://feeds.breakingnews.ie/bnireland"},
	}},

	{Name: "Italy", Flag: "🇮🇹", Sources: []Source{
		S{"https://www.thelocal.it/feeds/rss.php"},
	}},

	{Name: "Kazakhstan", Flag: "🇰🇿", Sources: []Source{
		S{"https://bnews.kz/en/rss/news"},
	}},

	{Name: "Kosovo", Flag: "🇽🇰", Sources: []Source{
		S{"https://www.theguardian.com/world/kosovo/rss"},
	}},

	{Name: "Latvia", Flag: "🇱🇻", Sources: []Source{
		S{"https://www.theguardian.com/world/latvia/rss"},
	}},

	{Name: "Liechtenstein", Flag: "🇱🇮", Sources: []Source{
		S{"https://www.nytimes.com/svc/collections/v1/publish/http://www.nytimes.com/topic/destination/liechtenstein/rss.xml"},
	}},

	{Name: "Lithuania", Flag: "🇱🇹", Sources: []Source{
		S{"https://www.theguardian.com/world/lithuania/rss"},
	}},

	{Name: "Luxembourg", Flag: "🇱🇺", Sources: []Source{
		S{"https://www.wort.lu/rss/en"},
	}},

	{Name: "Macedonia", Flag: "🇲🇰", Sources: []Source{
		S{"http://kurir.mk/en/?feed=atom"},
		S{"http://www.mia.mk/en/RSSFeed/FeedEN"},
	}},

	{Name: "Malta", Flag: "🇲🇹", Sources: []Source{
		S{"http://www.maltatoday.com.mt/rss/news/"},
	}},

	{Name: "Moldova", Flag: "🇲🇩", Sources: []Source{
		S{"http://en.publika.md/rss/"},
	}},

	{Name: "Monaco", Flag: "🇲🇨", Sources: []Source{
		S{"http://www.monacolife.net/feed/"},
	}},

	{Name: "Montenegro", Flag: "🇲🇪", Sources: []Source{
		S{"http://www.dailynewsmontenegro.com/feed"},
	}},

	{Name: "Netherlands", Flag: "🇳🇱", Sources: []Source{
		S{"https://www.dutchnews.nl/feed/?news"},
	}},

	{Name: "Norway", Flag: "🇳🇴", Sources: []Source{
		S{"https://www.thelocal.no/feeds/rss.php"},
	}},

	{Name: "Papua New Guinea", Flag: "🇵🇬", Sources: []Source{
		S{"http://www.thenational.com.pg/feed/"},
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

	{Name: "San Marino", Flag: "🇸🇲", Sources: []Source{
		S{"https://www.theguardian.com/world/san-marino/rss"},
	}},

	{Name: "Serbia", Flag: "🇷🇸", Sources: []Source{
		S{"https://inserbia.info/today/feed/"},
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

	{Name: "Turkey", Flag: "🇹🇷", Sources: []Source{
		S{"https://www.dailysabah.com/rss/turkey"},
	}},

	{Name: "Ukraine", Flag: "🇺🇦", Sources: []Source{
		S{"https://www.kyivpost.com/feed"},
	}},

	{Name: "Vatican City", Flag: "🇻🇦", Sources: []Source{
		S{"http://www.news.va/en/rss.xml"},
	}},
}
