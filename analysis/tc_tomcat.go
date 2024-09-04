package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var Tomcat = TC{
	Name:        "Customer Tomcat Legacy - shoud never fail",
	Application: data.CustomerTomcatLegacy,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
		},
	},
	Analysis: api.Analysis{
		Effort: 2,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:     "/addon/source/example-applications/example-1/src/main/resources/persistence.properties",
						Line:     2,
						Message:  "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
						CodeSnip: "jdbc.url=jdbc:oracle:thin:@169.60.225.216:1521/XEPDB1",
					},
					{
						File:     "/addon/source/example-applications/example-1/target/classes/persistence.properties",
						Line:     2,
						Message:  "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
						CodeSnip: "jdbc.url=jdbc:oracle:thin:@169.60.225.216:1521/XEPDB1",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Provider: "java",
				Name:     "io.konveyor.demo.config-utils",
				Version:  "1.0.0",
				Labels: []string{
					"konveyor.io/dep-source=internal",
					"konveyor.io/language=java",
				},
				SHA: "4010193B2F96CC7B7056C4CD51C3188FBBBC86E0",
			},
			{
				Provider: "java",
				Name:     "org.postgresql.postgresql",
				Version:  "42.2.23",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "cc8565ec39dbfee32c2c87f125162fe8a3010c28",
			},

			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-starter-actuator",
				Version:  "2.5.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "76dd6dea415751e05491337b7ff22bd08ae70c7e",
			},

			{
				Provider: "java",
				Name:     "com.oracle.database.jdbc.ojdbc8",
				Version:  "21.1.0.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "dea0cca54c29d3e44167cd80839692b325ae2daf",
			},
			{
				Provider: "java",
				Name:     "org.springframework.data.spring-data-jpa",
				Version:  "2.5.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "461ebcc9fc00dca10a754b0e96583ce7d281d312",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.hibernate-entitymanager",
				Version:  "5.4.32.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "b315696800e16d33bfb297d66f87a792caa3facc",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.12.3",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "ef6abf067337134089d074f411306a51f11a4d62",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-jdbc",
				Version:  "5.3.7",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "a4f87a03116ecde96213642141eb95da05022f51",
			},
			{Provider: "java",
				Name:    "org.springframework.spring-web",
				Version: "5.3.7",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "d9f78e0b045d90dc862cd4a39294a468b3cc6ba9",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-webmvc",
				Version:  "5.3.7",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "d0f042bff56bb90beabc6ed5d062fb87c69e652a",
			},
			{
				Provider: "java",
				Name:     "ch.qos.logback.logback-classic",
				Version:  "1.1.7",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "044c01db0f7d7aac366fb952a89c10251ed86f44",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.2.0.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "7f1beda5229a0c99a175603c18b3c66da44f966e",
			},
			{
				Provider: "java",
				Name:     "org.apache.tomcat.tomcat-servlet-api",
				Version:  "9.0.46",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "1f5ec6292bbca9e6c35172044b5fee0b0a97ef24",
			},
			{
				Provider: "java",
				Name:     "org.apache.tomcat.tomcat-jdbc",
				Version:  "9.0.46",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c3b975aba8359ecf35f6fca175c2e843a1d3c107",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.12.3",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "2b186d9cc73cfb9272171357d17f0979eac44889",
			},
			{
				Provider: "java",
				Name:     "javax.persistence.javax.persistence-api",
				Version:  "2.2",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "ac7080de51fc0596317c15e12ed441f7c0a84d09",
			},
			{
				Provider: "java",
				Name:     "com.sun.xml.fastinfoset.FastInfoset",
				Version:  "1.2.15",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "945cf1f4467c72add88309fb05cdf5e340b569f9",
			},
			{
				Provider: "java",
				Name:     "javax.xml.bind.jaxb-api",
				Version:  "2.3.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c42c51ae84892b73ef7de5351188908e673f5c69",
			},
			{
				Provider: "java",
				Name:     "net.bytebuddy.byte-buddy",
				Version:  "1.10.22",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "14de25cfee49cd27ae19153674bbb34c04c45d52",
			},
			{
				Provider: "java",
				Name:     "org.apache.logging.log4j.log4j-api",
				Version:  "2.14.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "9199a73770616b1ca0b00f576db3231aaab4876a",
			},
			{
				Provider: "java",
				Name:     "org.apache.logging.log4j.log4j-to-slf4j",
				Version:  "2.14.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "4638502177d694ad6f429a122e32f84ceba7db41",
			},
			{
				Provider: "java",
				Name:     "javax.activation.javax.activation-api",
				Version:  "1.2.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "1aa9ef58e50ba6868b2e955d61fcd73be5b4cea5",
			},
			{
				Provider: "java",
				Name:     "org.apache.tomcat.tomcat-juli",
				Version:  "9.0.46",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "1596051131c8426ebf744e0effed0e0005c87d57",
			},
			{
				Provider: "java",
				Name:     "jakarta.validation.jakarta.validation-api",
				Version:  "2.0.2",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "fc029778f5494ed05e5833f8bdb57e36dbda38aa",
			},
			{
				Provider: "java",
				Name:     "org.aspectj.aspectjrt",
				Version:  "1.9.6",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "2c4216b8c0f62edf69ec5cdd68619ba2aac5a4a1",
			},
			{
				Provider: "java",
				Name:     "org.checkerframework.checker-qual",
				Version:  "3.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "408a4451ff5bdef60400a49657867db100ea0f83",
			},
			{
				Provider: "java",
				Name:     "org.dom4j.dom4j",
				Version:  "2.1.3",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "012854caa63db09d82bf973bc37d7226aaaef463",
			},
			{
				Provider: "java",
				Name:     "org.glassfish.jaxb.jaxb-runtime",
				Version:  "2.3.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "1856da23a80b9b1374d925d6dcb4a21db2144204",
			},
			{
				Provider: "java",
				Name:     "org.glassfish.jaxb.txw2",
				Version:  "2.3.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c78aa440484eab1a6e2104e4fe69d0945a3cb3da",
			},
			{
				Provider: "java",
				Name:     "org.hdrhistogram.HdrHistogram",
				Version:  "2.1.12",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "9797702ee3e52e4be6bfbbc9fd20ac5447e7a541",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.common.hibernate-commons-annotations",
				Version:  "5.1.2.Final",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "573f22ce360cd7a8bcc0dae4deecbe4e8861007d",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.hibernate-core",
				Version:  "5.4.32.Final",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "5be381f7b6f3d4f17ce746e4ff54f4b8cdce40e4",
			},
			{
				Provider: "java",
				Name:     "jakarta.annotation.jakarta.annotation-api",
				Version:  "1.3.5",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "beb7649988a22ea30a17fcaeba8584323e86df74",
			},
			{
				Provider: "java",
				Name:     "io.micrometer.micrometer-core",
				Version:  "1.7.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "fd50ef746ed294d4e064c0cd3a14ca08543d139c",
			},
			{
				Provider: "java",
				Name:     "org.javassist.javassist",
				Version:  "3.27.0-GA",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0b7565662bc91e9648aab437135f32beb040ac15",
			},
			{
				Provider: "java",
				Name:     "org.jboss.jandex",
				Version:  "2.2.3.Final",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c70053a1326428ec641be311ccf5551a8ec76a63",
			},
			{
				Provider: "java",
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.4.1.Final",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "9d82f8eea1b5ed484775517d7588e320f9f7797a",
			},
			{
				Provider: "java",
				Name:     "org.jboss.spec.javax.transaction.jboss-transaction-api_1.2_spec",
				Version:  "1.1.1.Final",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "90823b310c573492696ad7e299b694ca2e70b4c1",
			},
			{
				Provider: "java",
				Name:     "org.jvnet.staxex.stax-ex",
				Version:  "1.8",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "cc7022b896125220e51f46fa50f4b68e564ffec1",
			},
			{
				Provider: "java",
				Name:     "org.latencyutils.LatencyUtils",
				Version:  "2.0.3",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "5baec26b6f9e5b17fdd200fc20af85eead4287c4",
			},
			{
				Provider: "java",
				Name:     "antlr.antlr",
				Version:  "2.7.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "52f15b99911ab8b8bc8744675f5cf1994a626fb8",
			},
			{
				Provider: "java",
				Name:     "org.slf4j.jul-to-slf4j",
				Version:  "1.7.30",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "f09448bdaeee63bc0644abae571b2d17c83d16c1",
			},
			{
				Provider: "java",
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.7.26",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "4d3419a58d77c07f49185aaa556a787d50508d27",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "48a6c425a45395e1ccfd99fd815c92d069040e43",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-actuator",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "ee202daac01b6399b857d187cfdbf6d97d6adc8f",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-actuator-autoconfigure",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c527193b5cc67f7534c27860171e44187746aaf5",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-autoconfigure",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "da542216009c858c2e8b32cb595578acc19d2df3",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-starter",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "391cbf83221ae09c1c0a471b25ab3221dfe46ef1",
			},
			{
				Provider: "java",
				Name:     "com.sun.istack.istack-commons-runtime",
				Version:  "3.0.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "8eb4c6b0e9b0a1fadf53fce8b3fc8415b00469ef",
			},
			{
				Provider: "java",
				Name:     "org.springframework.boot.spring-boot-starter-logging",
				Version:  "2.5.0",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "60f06908ef3b39d8c8780898e749c4c846fabb84",
			},
			{
				Provider: "java",
				Name:     "org.springframework.data.spring-data-commons",
				Version:  "2.5.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "bceeabb4ef399ba7ff8511f2931e1924a41cc921",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.datatype.jackson-datatype-jsr310",
				Version:  "2.12.3",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "db7822a553c167e95bdda25d0d6db44bd3abf847",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-aop",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0bf1d9d12108b8ab2d9d71d5fd5fee02d3ee5bde",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-beans",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "654397f55cd4a4734f8b76282e98c88884d0367a",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-context",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "67e3176098c81702c76d20977deec8101b3faf8c",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-core",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "44ce199d05bb1ce9682621cd18953ea307485fc1",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-expression",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "30bd0b3e802e5ba4e4d9fc68e57cc0e755ba9f9f",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-jcl",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "e1e7c14c73ae5fc616bb941ce8c1e7e62736cadf",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.12.3",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "87859f29ceebfab7a873c3b4f4b89c9a594b2842",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-orm",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "cc6911f3194cb77d493aa626c661789926027446",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-tx",
				Version:  "5.3.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c6df78e1d9b50b7063e4a196127d75ee9321f68b",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.classmate",
				Version:  "1.5.1",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "d5d564526c142037daead331ee5278c088777858",
			},
			{
				Provider: "java",
				Name:     "ch.qos.logback.logback-core",
				Version:  "1.1.7",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "6d1bdb1e28c56a8f989366b339f0f62545696e6d",
			},
			{
				Provider: "java",
				Name:     "org.yaml.snakeyaml",
				Version:  "1.28",
				Indirect: true,
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "3e38757e3eaf549cccd9bbdfa74b2930c177b8af",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
	},
}
