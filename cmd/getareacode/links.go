package main

type Source struct {
	ID    string
	Title string
	Url   string
}

var sources = []*Source{
	&Source{`2018`, `2018年6月中华人民共和国县以上行政区划代码`, `http://www.mca.gov.cn/article/sj/xzqh/2018/201804-12/20180608021501.html`},
	&Source{`2017`, `2017年中华人民共和国行政区划代码（截止2017年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/2018/201803131439.html`},
	&Source{`2016`, `2016年中华人民共和国行政区划代码（截止2016年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/2016/201612/201705311652.html`},
	&Source{`2015`, `2015年中华人民共和国行政区划代码（截止2015年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/2015/201706011127.html`},
	&Source{`2014`, `2014年中华人民共和国行政区划代码（截止2014年12月31日）`, `http://files2.mca.gov.cn/cws/201502/20150225163817214.html`},
	&Source{`2013`, `2013年中华人民共和国行政区划代码（截止2013年12月31日）`, `http://files2.mca.gov.cn/cws/201404/20140404125552372.htm`},
	&Source{`2012`, `2012年中华人民共和国行政区划代码（截止2012年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201707271556.html`},
	&Source{`2011`, `2011年中华人民共和国行政区划代码（截止2011年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201707271552.html`},
	&Source{`2010`, `2010年中华人民共和国行政区划代码（截止2010年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220946.html`},
	&Source{`2009`, `2009年中华人民共和国行政区划代码（截止2009年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220943.html`},
	&Source{`2008`, `2008年中华人民共和国行政区划代码（截止2008年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220941.html`},
	&Source{`2007`, `2007年中华人民共和国行政区划代码（截止2007年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220939.html`},
	&Source{`2006`, `2006年中华人民共和国行政区划代码（截止2006年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220936.html`},
	&Source{`2005`, `2005年中华人民共和国行政区划代码（截止2005年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220935.html`},
	&Source{`2004`, `2004年中华人民共和国行政区划代码（截止2004年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220930.html`},
	&Source{`2003`, `2003年中华人民共和国行政区划代码（截止2003年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220928.html`},
	&Source{`2002`, `2002年中华人民共和国行政区划代码（截止2002年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220927.html`},
	&Source{`2001`, `2001年中华人民共和国行政区划代码（截止2001年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220925.html`},
	&Source{`2000`, `2000年中华人民共和国行政区划代码（截止2000年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220923.html`},
	&Source{`1999`, `1999年中华人民共和国行政区划代码（截止1999年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220921.html`},
	&Source{`1998`, `1998年中华人民共和国行政区划代码（截止1998年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220918.html`},
	&Source{`1997`, `1997年中华人民共和国行政区划代码（截止1997年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220916.html`},
	&Source{`1996`, `1996年中华人民共和国行政区划代码（截止1996年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220914.html`},
	&Source{`1995`, `1995年中华人民共和国行政区划代码（截止1995年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220913.html`},
	&Source{`1994`, `1994年中华人民共和国行政区划代码（截止1994年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220911.html`},
	&Source{`1993`, `1993年中华人民共和国行政区划代码（截止1993年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041023.html`},
	&Source{`1992`, `1992年中华人民共和国行政区划代码（截止1992年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220910.html`},
	&Source{`1991`, `1991年中华人民共和国行政区划代码（截止1991年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041020.html`},
	&Source{`1990`, `1990年中华人民共和国行政区划代码（截止1990年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041018.html`},
	&Source{`1989`, `1989年中华人民共和国行政区划代码（截止1989年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041017.html`},
	&Source{`1988`, `1988年中华人民共和国行政区划代码（截止1988年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220903.html`},
	&Source{`1987`, `1987年中华人民共和国行政区划代码（截止1987年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220902.html`},
	&Source{`1986`, `1986年中华人民共和国行政区划代码（截止1986年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220859.html`},
	&Source{`1985`, `1985年中华人民共和国行政区划代码（截止1985年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220858.html`},
	&Source{`1984`, `1984年中华人民共和国行政区划代码（截止1984年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220856.html`},
	&Source{`1983`, `1983年中华人民共和国行政区划代码（截止1983年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708160821.html`},
	&Source{`1982`, `1982年中华人民共和国行政区划代码（截止1982年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/1980-2000/201707141125.html`},
	&Source{`1981`, `1981年中华人民共和国行政区划代码（截止1981年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041004.html`},
	&Source{`1980`, `1980年中华人民共和国行政区划代码（截止1980年12月31日）`, `http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708040959.html`},
}
