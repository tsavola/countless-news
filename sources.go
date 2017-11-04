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

	{Name: "Afghanistan", Flag: "🇦🇫", Sources: []Source{
		S{"https://www.pajhwok.com/en/nodequeue/1/feed"},
	}},

	{Name: "Albania", Flag: "🇦🇱", Sources: []Source{
		S{"http://www.tiranatimes.com/?feed=rss2"},
	}},

	{Name: "Andorra", Flag: "🇦🇩", Sources: []Source{
		S{"https://all-andorra.com/feed/"},
	}},

	{Name: "Armenia", Flag: "🇦🇲", Sources: []Source{
		S{"http://arka.am/rss_news.php?lang=en"},
	}},

	{Name: "Australia", Flag: "🇦🇺", Sources: []Source{
		S{"http://www.news.com.au/feed/"},
	}},

	{Name: "Austria", Flag: "🇦🇹", Sources: []Source{
		S{"https://www.thelocal.at/feeds/rss.php"},
	}},

	{Name: "Azerbaijan", Flag: "🇦🇿", Sources: []Source{
		S{"http://www.today.az/rss.php"},
	}},

	{Name: "Bahrain", Flag: "🇧🇭", Sources: []Source{
		S{"http://www.twentyfoursevennews.com/feed/"},
	}},

	{Name: "Bangladesh", Flag: "🇧🇩", Sources: []Source{
		S{"https://bdnews24.com/?widgetName=rssfeed&widgetId=1150&getXmlFeed=true"},
	}},

	{Name: "Belarus", Flag: "🇧🇾", Sources: []Source{
		S{"http://eng.belta.by/rss"},
	}},

	{Name: "Belgium", Flag: "🇧🇪", Sources: []Source{
		S{"https://deredactie.be/cm/vrtnieuws.english?mode=atom"},
	}},

	{Name: "Bhutan", Flag: "🇧🇹", Sources: []Source{
		S{"http://www.kuenselonline.com/feed/"},
	}},

	{Name: "Bosnia and Herzegovina", Flag: "🇧🇦", Sources: []Source{
		S{"http://www.telegraph.co.uk/news/rss.xml"},
		S{"https://www.nytimes.com/svc/collections/v1/publish/http://www.nytimes.com/topic/destination/bosnia-and-herzegovina/rss.xml"},
		S{"https://www.theguardian.com/world/bosnia-and-herzegovina/rss"},
	}},

	{Name: "Bulgaria", Flag: "🇧🇬", Sources: []Source{
		S{"https://feeds.feedburner.com/TheSofiaGlobe"},
	}},

	{Name: "Brunei", Flag: "🇧🇳", Sources: []Source{
		S{"http://www.theborneopost.com/news/brunei/feed/"},
	}},

	{Name: "Cambodia", Flag: "🇰🇭", Sources: []Source{
		S{"https://www.cambodiadaily.com/feed/"},
	}},

	{Name: "China", Flag: "🇨🇳", Sources: []Source{
		S{"https://www.ecns.cn/rss/rss.xml"},
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

	{Name: "Egypt", Flag: "🇪🇬", Sources: []Source{
		S{"http://www.egyptindependent.com/category/egypt/feed/"},
	}},

	{Name: "Estonia", Flag: "🇪🇪", Sources: []Source{
		S{"https://news.postimees.ee/rss"},
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

	{Name: "India", Flag: "🇮🇳", Sources: []Source{
		S{"http://indianexpress.com/section/india/feed/"},
		S{"https://timesofindia.indiatimes.com/rssfeeds/-2128936835.cms"},
	}},

	{Name: "Indonesia", Flag: "🇮🇩", Sources: []Source{
		S{"http://www.abc.net.au/news/feed/26574/rss.xml"},
	}},

	{Name: "Iran", Flag: "🇮🇷", Sources: []Source{
		S{"http://www.iran-daily.com/News/Feed?id=3"},
		S{"http://www.payvand.com/news/rssfeed.xml"},
	}},

	{Name: "Iraq", Flag: "🇮🇶", Sources: []Source{
		S{"https://www.theguardian.com/world/iraq/rss"},
	}},

	{Name: "Ireland", Flag: "🇮🇪", Sources: []Source{
		S{"http://feeds.breakingnews.ie/bnireland"},
	}},

	{Name: "Israel", Flag: "🇮🇱", Sources: []Source{
		S{"https://www.haaretz.com/cmlink/1.628764"},
		S{"https://www.jpost.com/Rss/RssFeedsIsraelNews.aspx"},
	}},

	{Name: "Italy", Flag: "🇮🇹", Sources: []Source{
		S{"https://www.thelocal.it/feeds/rss.php"},
	}},

	{Name: "Japan", Flag: "🇯🇵", Sources: []Source{
		S{"https://japantoday.com/feed/atom"},
		S{"https://www.japantimes.co.jp/news/national/feed/"},
	}},

	{Name: "Jordan", Flag: "🇯🇴", Sources: []Source{
		S{"https://www.jordantimes.com/rss.xml"},
	}},

	{Name: "Kazakhstan", Flag: "🇰🇿", Sources: []Source{
		S{"https://bnews.kz/en/rss/news"},
	}},

	{Name: "Kosovo", Flag: "🇽🇰", Sources: []Source{
		S{"https://www.theguardian.com/world/kosovo/rss"},
	}},

	{Name: "Kuwait", Flag: "🇰🇼", Sources: []Source{
		S{"http://news.kuwaittimes.net/website/feed/"},
	}},

	{Name: "Kyrgyzstan", Flag: "🇰🇬", Sources: []Source{
		S{"http://kabar.kg/eng/rss/"},
	}},

	{Name: "Laos", Flag: "🇱🇦", Sources: []Source{
		S{"https://laotiantimes.com/feed/"},
	}},

	{Name: "Latvia", Flag: "🇱🇻", Sources: []Source{
		S{"https://www.theguardian.com/world/latvia/rss"},
	}},

	{Name: "Lebanon", Flag: "🇱🇧", Sources: []Source{
		S{"https://www.dailystar.com.lb/RSS.aspx?live=1"},
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

	{Name: "Malaysia", Flag: "🇲🇾", Sources: []Source{
		S{"https://www.malaysiakini.com/en/news.rss"},
	}},

	{Name: "Maldives", Flag: "🇲🇻", Sources: []Source{
		S{"http://en.mihaaru.com/feed/"},
		S{"https://maldivesindependent.com/feed"},
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

	{Name: "Mongolia", Flag: "🇲🇳", Sources: []Source{
		S{"http://mongolia.gogo.mn/feed"},
		S{"https://theubpost.mn/feed/"},
	}},

	{Name: "Montenegro", Flag: "🇲🇪", Sources: []Source{
		S{"http://www.dailynewsmontenegro.com/feed"},
	}},

	{Name: "Myanmar", Flag: "🇲🇲", Sources: []Source{
		S{"https://www.mizzima.com/rss.xml"},
	}},

	{Name: "Nepal", Flag: "🇳🇵", Sources: []Source{
		S{"http://english.onlinekhabar.com/feed"},
		S{"https://thehimalayantimes.com/category/nepal/feed/"},
	}},

	{Name: "Netherlands", Flag: "🇳🇱", Sources: []Source{
		S{"https://www.dutchnews.nl/feed/?news"},
	}},

	{Name: "North Korea", Flag: "🇰🇵", Sources: []Source{
		S{"https://www.koreatimes.co.kr/www/rss/northkorea.xml"},
		S{"https://www.nknews.org/feed/"},
	}},

	{Name: "Norway", Flag: "🇳🇴", Sources: []Source{
		S{"https://www.thelocal.no/feeds/rss.php"},
	}},

	{Name: "Oman", Flag: "🇴🇲", Sources: []Source{
		S{"http://www.muscatdaily.com/rss/feed/Muscat_Daily_Oman_News"},
	}},

	{Name: "Pakistan", Flag: "🇵🇰", Sources: []Source{
		S{"http://dunyanews.tv/news.xml"},
		S{"https://tribune.com.pk/pakistan/feed/"},
	}},

	{Name: "Palestine", Flag: "🇵🇸", Sources: []Source{
		S{"http://english.pnn.ps/feed/"},
	}},

	{Name: "Papua New Guinea", Flag: "🇵🇬", Sources: []Source{
		S{"http://www.thenational.com.pg/feed/"},
	}},

	{Name: "Philippines", Flag: "🇵🇭", Sources: []Source{
		S{"http://www.philstar.com/rss/nation"},
		S{"https://www.inquirer.net/fullfeed"},
	}},

	{Name: "Poland", Flag: "🇵🇱", Sources: []Source{
		S{"http://www.thenews.pl/Rss/47e92f6a-fbb8-4bb3-9b2e-128125c4b722"},
	}},

	{Name: "Portugal", Flag: "🇵🇹", Sources: []Source{
		S{"https://portugalresident.com/articles.xml"},
	}},

	{Name: "Qatar", Flag: "🇶🇦", Sources: []Source{
		S{"https://dohanews.co/feed/"},
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

	{Name: "Saudi Arabia", Flag: "🇸🇦", Sources: []Source{
		S{"http://www.arabnews.com/cat/1/rss.xml"},
		S{"https://saudigazette.com.sa/rssFeed/72"},
	}},

	{Name: "Serbia", Flag: "🇷🇸", Sources: []Source{
		S{"https://inserbia.info/today/feed/"},
	}},

	{Name: "Singapore", Flag: "🇸🇬", Sources: []Source{
		S{"http://www.straitstimes.com/news/singapore/rss.xml"},
		S{"http://www.theindependent.sg/feed/"},
		S{"https://www.channelnewsasia.com/rssfeeds/8396082"},
	}},

	{Name: "Slovakia", Flag: "🇸🇰", Sources: []Source{
		S{"https://spectator.sme.sk/rss-title"},
	}},

	{Name: "Slovenia", Flag: "🇸🇮", Sources: []Source{
		S{"http://www.sloveniatimes.com/rss?category_id=1"},
	}},

	{Name: "South Korea", Flag: "🇰🇷", Sources: []Source{
		S{"http://www.koreaherald.com/rss_xml.php"},
		S{"https://www.koreatimes.co.kr/www/rss/nation.xml"},
	}},

	{Name: "Spain", Flag: "🇪🇸", Sources: []Source{
		S{"https://www.thelocal.es/feeds/rss.php"},
	}},

	{Name: "Sri Lanka", Flag: "🇱🇰", Sources: []Source{
		S{"http://www.dailymirror.lk/RSS_Feeds/breaking-news"},
		S{"https://www.adaderana.lk/rss.php"},
		S{"https://www.dailynews.lk/taxonomy/term/799/all/feed"},
	}},

	{Name: "Sweden", Flag: "🇸🇪", Sources: []Source{
		S{"https://www.thelocal.se/feeds/rss.php"},
	}},

	{Name: "Switzerland", Flag: "🇨🇭", Sources: []Source{
		S{"https://www.thelocal.ch/feeds/rss.php"},
	}},

	{Name: "Syria", Flag: "🇸🇾", Sources: []Source{
		S{"http://sana.sy/en/?feed=rss2"},
	}},

	{Name: "United Kingdom", Flag: "🇬🇧", Sources: []Source{
		S{"https://feeds.bbci.co.uk/news/uk/rss.xml"},
		S{"https://www.independent.co.uk/news/uk/rss"},
		S{"https://www.theguardian.com/uk-news/rss"},
	}},

	{Name: "Taiwan", Flag: "🇹🇼", Sources: []Source{
		S{"http://api.taiwantoday.tw/en/rss.php?unit=2,6,10,15,18"},
	}},

	{Name: "Tajikistan", Flag: "🇹🇯", Sources: []Source{
		S{"https://www.theguardian.com/world/tajikistan/rss"},
	}},

	{Name: "Thailand", Flag: "🇹🇭", Sources: []Source{
		S{"https://pattayaone.news/en/category/pattayaone-new-en/national/feed/"},
		S{"https://www.bangkokpost.com/rss/data/topstories.xml"},
		S{"https://www.thephuketnews.com/rss-xml/news-thailand.xml"},
	}},

	{Name: "Turkey", Flag: "🇹🇷", Sources: []Source{
		S{"https://www.dailysabah.com/rss/turkey"},
	}},

	{Name: "Turkmenistan", Flag: "🇹🇲", Sources: []Source{
		S{"http://www.turkmenistan.ru/en/taxonomy/term/3/all/feed"},
	}},

	{Name: "Ukraine", Flag: "🇺🇦", Sources: []Source{
		S{"https://www.kyivpost.com/feed"},
	}},

	{Name: "United Arab Emirates", Flag: "🇦🇪", Sources: []Source{
		S{"https://www.khaleejtimes.com/rss.xml"},
	}},

	{Name: "Uzbekistan", Flag: "🇺🇿", Sources: []Source{
		S{"http://uza.uz/en/rss/"},
		S{"http://www.ut.uz/en/rss/"},
		S{"https://www.uzdaily.com/rss.htm"},
	}},

	{Name: "Vatican City", Flag: "🇻🇦", Sources: []Source{
		S{"http://www.news.va/en/rss.xml"},
	}},

	{Name: "Vietnam", Flag: "🇻🇳", Sources: []Source{
		S{"http://vietnamnews.vn/rss/ovietnam.rss"},
	}},

	{Name: "Yemen", Flag: "🇾🇪", Sources: []Source{
		S{"https://www.almasdarnews.com/article/category/yemen/feed/"},
	}},
}
