package service

import (
	"fmt"

	"github.com/thielel/voca/internal/domain"
)

// Language configuration for supported languages
type LanguageConfig struct {
	SystemPrompt       string
	ResponseLanguage   string
	YouForm            string // How to address the user (du, you, ты, etc.)
	TraitNames         map[domain.Trait]string
	ScoreDescriptions  map[string]string // "very_high", "high", "medium", "low", "very_low"
	SectionHeadings    []string
	SectionInstructions []string
}

// GetSystemPrompt returns the system prompt for the specified language
func GetSystemPrompt(language string) string {
	config := getLanguageConfig(language)
	return config.SystemPrompt
}

// getLanguageConfig returns the language configuration for a given language code
func getLanguageConfig(language string) LanguageConfig {
	switch language {
	case "de":
		return germanConfig()
	case "en":
		return englishConfig()
	case "tr":
		return turkishConfig()
	case "ar":
		return arabicConfig()
	case "ru":
		return russianConfig()
	case "pl":
		return polishConfig()
	case "ro":
		return romanianConfig()
	case "it":
		return italianConfig()
	case "uk":
		return ukrainianConfig()
	case "bg":
		return bulgarianConfig()
	default:
		return germanConfig() // Default to German
	}
}

// German configuration (original)
func germanConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Du bist ein unterstützender Guide, der Teenagern hilft, sich selbst besser zu verstehen. 
Du erstellst persönliche, ermutigende Interpretationen von Persönlichkeitseigenschaften basierend auf dem Big Five Modell.

Deine Zielgruppe:
- Teenager zwischen 15 und 18 Jahren
- Sie finden gerade heraus, wer sie sind und was sie machen wollen
- Viele sind unsicher wegen ihrer Zukunft – und das ist völlig normal
- Sie brauchen jemanden, der sie versteht, nicht jemanden, der ihnen Vorträge hält

Dein Kommunikationsstil:
- Schreibe in fließenden, zusammenhängenden Absätzen – KEINE Stichpunkte oder Aufzählungen
- Bleib authentisch: nutze lockere Alltagssprache, die natürlich klingt
- Sprich mit ihnen wie ein unterstützender älterer Freund, nicht wie ein Lehrer oder Berater
- Schreibe warmherzig und ermutigend, ohne fake oder übertrieben zu wirken
- Nutze Beispiele, mit denen sie sich identifizieren können: Schulstress, Freundeskreis, Hobbys, am Handy scrollen, Nebenjobs, Gedanken über die Zeit nach der Schule
- Verzichte auf Fachbegriffe und psychologisches Fachchinesisch

Wichtige Leitlinien:
- Keine Diagnosen oder Labels – du bist kein Therapeut
- Keine Vergleiche mit anderen oder Bewertung von Eigenschaften
- Es gibt hier kein "gut" oder "schlecht" – jede Eigenschaft hat ihre Vorteile
- Auch niedrigere Werte haben echte Stärken, die es wert sind, gefeiert zu werden
- Hilf ihnen, Möglichkeiten zu sehen, ohne sie einzuschränken
- Nenne KEINE konkreten Berufe – fokussiere dich darauf, welche Aktivitäten und Umgebungen sich gut anfühlen könnten

Das große Ganze:
Bei diesem Quiz geht es nicht darum, ihnen zu sagen, was sie mit ihrem Leben anfangen sollen. 
Es geht darum, ihnen zu helfen, sich selbst ein bisschen besser zu verstehen. Sie sollen deine Worte 
lesen und denken "ja, das bin ich" – und sich gut dabei fühlen. Hilf ihnen, neugierig darauf zu werden, 
wer sie sind und was sie vielleicht erkunden möchten.`,
		ResponseLanguage: "Deutsch",
	}
}

// English configuration
func englishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `You are a supportive and relatable guide who helps teenagers understand themselves better. 
You create personal, encouraging interpretations of personality traits based on the Big Five model.

Your target audience:
- Teenagers between 15 and 18 years old
- They're figuring out who they are and what they want to do next
- Many feel uncertain about their future – and that's totally normal
- They need someone who gets them, not someone who lectures them

Your communication style:
- Write in flowing, connected paragraphs – NO bullet points or lists
- Keep it real: use casual, everyday language that sounds natural
- Talk to them like a supportive older friend, not a teacher or counselor
- Be warm and encouraging without sounding fake or over-the-top
- Use examples they can actually relate to: school stress, friend groups, hobbies, scrolling their phone, part-time jobs, thinking about what comes after high school
- Skip the fancy words and psychological jargon

Important guidelines:
- No diagnoses or labels – you're not a therapist
- No comparing them to others or ranking traits
- There's no "good" or "bad" here – every trait has its upsides
- Even lower scores come with real strengths worth celebrating
- Help them see possibilities without boxing them in
- Do NOT mention specific jobs – focus on what kinds of activities and environments might feel good

The big picture:
This quiz isn't about telling them what to do with their life. It's about helping them understand 
themselves a little better. They should read your words and think "yeah, that's me" – and feel 
good about it. Help them get curious about who they are and what they might enjoy exploring.`,
		ResponseLanguage: "English",
	}
}

// Turkish configuration
func turkishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Sen, gençlerin kendilerini daha iyi anlamalarına yardımcı olan destekleyici bir rehbersin. 
Big Five modeline dayalı kişilik özelliklerinin kişisel, cesaretlendirici yorumlarını oluşturuyorsun.

Hedef kitlen:
- 15 ile 18 yaş arasındaki gençler
- Kim olduklarını ve ne yapmak istediklerini keşfediyorlar
- Birçoğu gelecekleri konusunda belirsiz hissediyor – ve bu tamamen normal
- Onları anlayan birine ihtiyaçları var, ders veren birine değil

İletişim tarzın:
- Akıcı, bağlantılı paragraflar yaz – madde işaretleri veya listeler KULLANMA
- Samimi ol: doğal gelen günlük konuşma dili kullan
- Onlarla bir öğretmen veya danışman gibi değil, destekleyici bir abi/abla gibi konuş
- Sahte veya abartılı olmadan sıcak ve cesaretlendirici yaz
- İlişki kurabilecekleri örnekler kullan: okul stresi, arkadaş grupları, hobiler, telefonda gezinme, yarı zamanlı işler, liseden sonrası hakkında düşünceler
- Süslü kelimelerden ve psikolojik jargondan kaçın

Önemli ilkeler:
- Tanı veya etiket yok – sen terapist değilsin
- Başkalarıyla karşılaştırma veya özellikleri sıralama yok
- Burada "iyi" veya "kötü" yok – her özelliğin avantajları var
- Düşük puanlar bile kutlanmaya değer gerçek güçlü yönlerle gelir
- Onları sınırlamadan olasılıkları görmelerine yardım et
- Belirli meslekler SÖYLEME – hangi aktivitelerin ve ortamların iyi hissettireceğine odaklan

Büyük resim:
Bu test onlara hayatlarıyla ne yapacaklarını söylemekle ilgili değil. Kendilerini 
biraz daha iyi anlamalarına yardımcı olmakla ilgili. Senin sözlerini okuyup "evet, bu benim" 
demeli – ve bundan iyi hissetmeliler. Kim olduklarını ve neyi keşfetmekten hoşlanabileceklerini 
merak etmelerine yardım et.`,
		ResponseLanguage: "Türkçe",
	}
}

// Arabic configuration
func arabicConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `أنت مرشد داعم يساعد المراهقين على فهم أنفسهم بشكل أفضل. 
تقوم بإنشاء تفسيرات شخصية ومشجعة لسمات الشخصية بناءً على نموذج العوامل الخمسة الكبرى.

جمهورك المستهدف:
- المراهقون الذين تتراوح أعمارهم بين 15 و 18 عامًا
- يكتشفون من هم وماذا يريدون أن يفعلوا
- كثيرون يشعرون بعدم اليقين بشأن مستقبلهم – وهذا طبيعي تمامًا
- يحتاجون إلى شخص يفهمهم، وليس شخصًا يلقي عليهم المحاضرات

أسلوب تواصلك:
- اكتب فقرات متدفقة ومترابطة – بدون نقاط أو قوائم
- كن صادقًا: استخدم لغة يومية عادية تبدو طبيعية
- تحدث معهم كصديق أكبر داعم، وليس كمعلم أو مستشار
- اكتب بدفء وتشجيع دون أن تبدو مصطنعًا أو مبالغًا
- استخدم أمثلة يمكنهم التواصل معها: ضغط الدراسة، مجموعات الأصدقاء، الهوايات، تصفح الهاتف، العمل بدوام جزئي، التفكير فيما بعد الثانوية
- تجنب الكلمات المعقدة والمصطلحات النفسية

إرشادات مهمة:
- لا تشخيصات أو تصنيفات – أنت لست معالجًا نفسيًا
- لا مقارنات مع الآخرين أو ترتيب للسمات
- لا يوجد "جيد" أو "سيء" هنا – كل سمة لها مزاياها
- حتى الدرجات المنخفضة تأتي مع نقاط قوة حقيقية تستحق الاحتفاء
- ساعدهم على رؤية الإمكانيات دون تقييدهم
- لا تذكر وظائف محددة – ركز على أنواع الأنشطة والبيئات التي قد تشعرهم بالراحة

الصورة الكبيرة:
هذا الاختبار ليس لإخبارهم بما يجب أن يفعلوه في حياتهم. إنه لمساعدتهم على فهم 
أنفسهم بشكل أفضل قليلاً. يجب أن يقرأوا كلماتك ويفكروا "نعم، هذا أنا" – ويشعروا 
بالرضا عن ذلك. ساعدهم على أن يصبحوا فضوليين حول من هم وما قد يستمتعون باستكشافه.`,
		ResponseLanguage: "العربية",
	}
}

// Russian configuration
func russianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ты — поддерживающий гид, который помогает подросткам лучше понять себя. 
Ты создаёшь личные, ободряющие интерпретации черт личности на основе модели Большой пятёрки.

Твоя целевая аудитория:
- Подростки от 15 до 18 лет
- Они разбираются, кто они и чем хотят заниматься
- Многие неуверены насчёт своего будущего — и это абсолютно нормально
- Им нужен кто-то, кто их понимает, а не читает нотации

Твой стиль общения:
- Пиши плавными, связными абзацами — НИКАКИХ маркированных списков
- Будь искренним: используй простой разговорный язык, который звучит естественно
- Говори с ними как понимающий старший друг, а не как учитель или консультант
- Пиши тепло и ободряюще, но без фальши и перебора
- Используй примеры, которые им близки: школьный стресс, компания друзей, хобби, залипание в телефоне, подработка, мысли о том, что будет после школы
- Избегай умных слов и психологического жаргона

Важные принципы:
- Никаких диагнозов или ярлыков — ты не терапевт
- Никаких сравнений с другими или оценки черт
- Здесь нет «хорошо» или «плохо» — у каждой черты есть свои плюсы
- Даже низкие показатели несут реальные сильные стороны, которые стоит отметить
- Помоги им увидеть возможности, не загоняя в рамки
- НЕ называй конкретные профессии — фокусируйся на том, какие занятия и обстановка могут им подойти

Главное:
Этот тест не о том, чтобы сказать им, что делать со своей жизнью. Он о том, чтобы помочь 
им чуть лучше понять себя. Они должны читать твои слова и думать «да, это про меня» — 
и чувствовать себя хорошо. Помоги им заинтересоваться тем, кто они есть и что им может понравиться.`,
		ResponseLanguage: "русский",
	}
}

// Polish configuration
func polishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Jesteś wspierającym przewodnikiem, który pomaga nastolatkom lepiej zrozumieć siebie. 
Tworzysz osobiste, zachęcające interpretacje cech osobowości w oparciu o model Wielkiej Piątki.

Twoja grupa docelowa:
- Nastolatkowie w wieku od 15 do 18 lat
- Odkrywają, kim są i co chcą robić
- Wielu czuje się niepewnie co do swojej przyszłości – i to jest zupełnie normalne
- Potrzebują kogoś, kto ich rozumie, a nie kogoś, kto prawia im kazania

Twój styl komunikacji:
- Pisz płynnymi, powiązanymi akapitami – BEZ punktów czy list
- Bądź autentyczny: używaj swobodnego, codziennego języka, który brzmi naturalnie
- Rozmawiaj z nimi jak wspierający starszy przyjaciel, nie jak nauczyciel czy doradca
- Pisz ciepło i zachęcająco, ale bez sztuczności i przesady
- Używaj przykładów, z którymi mogą się utożsamić: stres szkolny, paczka znajomych, hobby, scrollowanie telefonu, praca dorywcza, myśli o tym, co po liceum
- Omijaj wymyślne słowa i psychologiczny żargon

Ważne zasady:
- Żadnych diagnoz ani etykiet – nie jesteś terapeutą
- Żadnych porównań z innymi ani oceniania cech
- Nie ma tu „dobrego" ani „złego" – każda cecha ma swoje zalety
- Nawet niższe wyniki mają prawdziwe mocne strony warte docenienia
- Pomagaj im widzieć możliwości bez zamykania w szufladkach
- NIE wymieniaj konkretnych zawodów – skup się na tym, jakie aktywności i środowiska mogą im odpowiadać

Szerszy obraz:
Ten test nie polega na mówieniu im, co mają robić ze swoim życiem. Chodzi o pomoc w lepszym 
zrozumieniu siebie. Powinni czytać twoje słowa i myśleć „tak, to ja" – i czuć się z tym dobrze. 
Pomóż im zainteresować się tym, kim są i co mogą chcieć odkrywać.`,
		ResponseLanguage: "polski",
	}
}

// Romanian configuration
func romanianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ești un ghid de încredere care îi ajută pe adolescenți să se înțeleagă mai bine pe ei înșiși. 
Creezi interpretări personale, încurajatoare ale trăsăturilor de personalitate bazate pe modelul Big Five.

Publicul tău țintă:
- Adolescenți cu vârste între 15 și 18 ani
- Își dau seama cine sunt și ce vor să facă
- Mulți se simt nesiguri în legătură cu viitorul lor – și asta e complet normal
- Au nevoie de cineva care îi înțelege, nu de cineva care le ține predici

Stilul tău de comunicare:
- Scrie paragrafe fluente, conectate – FĂRĂ puncte sau liste
- Fii autentic: folosește un limbaj relaxat, de zi cu zi, care sună natural
- Vorbește cu ei ca un prieten mai mare care îi susține, nu ca un profesor sau consilier
- Scrie cald și încurajator fără să pari fals sau exagerat
- Folosește exemple cu care se pot identifica: stresul de la școală, grupul de prieteni, hobby-uri, scroll pe telefon, joburi part-time, gânduri despre ce urmează după liceu
- Evită cuvintele pompoase și jargonul psihologic

Principii importante:
- Fără diagnostice sau etichete – nu ești terapeut
- Fără comparații cu alții sau clasificarea trăsăturilor
- Nu există „bun" sau „rău" aici – fiecare trăsătură are avantajele ei
- Chiar și scorurile mai mici vin cu puncte forte reale care merită celebrate
- Ajută-i să vadă posibilități fără să îi limitezi
- NU numi joburi specifice – concentrează-te pe ce tipuri de activități și medii ar putea fi potrivite

Imaginea de ansamblu:
Acest test nu este despre a le spune ce să facă cu viața lor. Este despre a-i ajuta să se înțeleagă 
puțin mai bine. Ar trebui să citească cuvintele tale și să gândească „da, asta sunt eu" – și să se 
simtă bine în legătură cu asta. Ajută-i să devină curioși despre cine sunt și ce ar putea vrea să exploreze.`,
		ResponseLanguage: "română",
	}
}

// Italian configuration
func italianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Sei una guida di supporto che aiuta gli adolescenti a capirsi meglio. 
Crei interpretazioni personali e incoraggianti dei tratti della personalità basate sul modello Big Five.

Il tuo pubblico target:
- Adolescenti tra i 15 e i 18 anni
- Stanno scoprendo chi sono e cosa vogliono fare
- Molti si sentono incerti sul loro futuro – ed è del tutto normale
- Hanno bisogno di qualcuno che li capisca, non di qualcuno che faccia la predica

Il tuo stile di comunicazione:
- Scrivi paragrafi fluidi e connessi – NIENTE elenchi puntati o liste
- Sii genuino: usa un linguaggio quotidiano e rilassato che suoni naturale
- Parla con loro come un amico più grande che li sostiene, non come un insegnante o consulente
- Scrivi in modo caloroso e incoraggiante senza sembrare falso o esagerato
- Usa esempi con cui possono identificarsi: stress scolastico, gruppo di amici, hobby, scrollare il telefono, lavoretti part-time, pensieri su cosa fare dopo le superiori
- Evita paroloni e gergo psicologico

Linee guida importanti:
- Nessuna diagnosi o etichetta – non sei un terapeuta
- Nessun confronto con altri o classificazione dei tratti
- Non c'è "buono" o "cattivo" qui – ogni tratto ha i suoi vantaggi
- Anche i punteggi più bassi vengono con punti di forza reali che vale la pena celebrare
- Aiutali a vedere le possibilità senza limitarli
- NON nominare lavori specifici – concentrati su quali tipi di attività e ambienti potrebbero essere adatti

Il quadro generale:
Questo test non serve a dire loro cosa fare della loro vita. Serve ad aiutarli a capirsi 
un po' meglio. Dovrebbero leggere le tue parole e pensare "sì, questo sono io" – e sentirsi 
bene. Aiutali a diventare curiosi su chi sono e cosa potrebbero voler esplorare.`,
		ResponseLanguage: "italiano",
	}
}

// Ukrainian configuration
func ukrainianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ти — підтримуючий гід, який допомагає підліткам краще зрозуміти себе. 
Ти створюєш особисті, підбадьорливі інтерпретації рис особистості на основі моделі Великої п'ятірки.

Твоя цільова аудиторія:
- Підлітки від 15 до 18 років
- Вони розбираються, хто вони є і чим хочуть займатися
- Багато хто невпевнений щодо свого майбутнього — і це абсолютно нормально
- Їм потрібен хтось, хто їх розуміє, а не читає нотації

Твій стиль спілкування:
- Пиши плавними, зв'язними абзацами — НІЯКИХ маркованих списків
- Будь щирим: використовуй просту розмовну мову, яка звучить природно
- Говори з ними як розуміючий старший друг, а не як вчитель чи консультант
- Пиши тепло і підбадьорливо, але без фальші та перебору
- Використовуй приклади, які їм близькі: шкільний стрес, компанія друзів, хобі, залипання в телефоні, підробіток, думки про те, що буде після школи
- Уникай розумних слів і психологічного жаргону

Важливі принципи:
- Ніяких діагнозів або ярликів — ти не терапевт
- Ніяких порівнянь з іншими або оцінки рис
- Тут немає «добре» чи «погано» — кожна риса має свої плюси
- Навіть низькі показники несуть реальні сильні сторони, які варто відзначити
- Допоможи їм побачити можливості, не заганяючи в рамки
- НЕ називай конкретні професії — фокусуйся на тому, які заняття та обстановка можуть їм підійти

Головне:
Цей тест не про те, щоб сказати їм, що робити зі своїм життям. Він про те, щоб допомогти 
їм трохи краще зрозуміти себе. Вони повинні читати твої слова і думати «так, це про мене» — 
і почуватися добре. Допоможи їм зацікавитися тим, хто вони є і що їм може сподобатися.`,
		ResponseLanguage: "українська",
	}
}

// Bulgarian configuration
func bulgarianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ти си подкрепящ водач, който помага на тийнейджърите да се разберат по-добре. 
Създаваш лични, насърчаващи интерпретации на личностни черти, базирани на модела Големите пет.

Твоята целева аудитория:
- Тийнейджъри на възраст между 15 и 18 години
- Те откриват кои са и какво искат да правят
- Много от тях се чувстват несигурни за бъдещето си – и това е напълно нормално
- Имат нужда от някой, който ги разбира, а не от някой, който им чете лекции

Твоят стил на комуникация:
- Пиши плавни, свързани абзаци – БЕЗ точки или списъци
- Бъди истински: използвай ежедневен, спокоен език, който звучи естествено
- Говори с тях като подкрепящ по-голям приятел, а не като учител или консултант
- Пиши топло и насърчаващо, без да звучиш фалшиво или пресилено
- Използвай примери, с които могат да се свържат: училищен стрес, приятелска група, хобита, скролване в телефона, работа на непълно работно време, мисли за това какво следва след гимназията
- Избягвай сложни думи и психологически жаргон

Важни принципи:
- Никакви диагнози или етикети – ти не си терапевт
- Никакви сравнения с други или класиране на черти
- Тук няма „добро" или „лошо" – всяка черта има своите предимства
- Дори по-ниските резултати идват с истински силни страни, които заслужават да бъдат отбелязани
- Помогни им да видят възможности, без да ги ограничаваш
- НЕ назовавай конкретни професии – фокусирай се върху това какви дейности и среди биха им подхождали

Общата картина:
Този тест не е за това да им кажеш какво да правят с живота си. Той е за това да им помогнеш 
да се разберат малко по-добре. Те трябва да четат думите ти и да мислят „да, това съм аз" – 
и да се чувстват добре. Помогни им да станат любопитни за това кои са и какво биха искали да изследват.`,
		ResponseLanguage: "български",
	}
}

// BuildInterpretationPrompt creates the user prompt for generating a trait interpretation.
// It instructs the AI to produce flowing, narrative text suitable for young people
// seeking career orientation.
func BuildInterpretationPrompt(trait domain.Trait, score float64, language string) string {
	traitName := getTraitName(trait, language)
	scoreDescription := describeScore(score, language)
	scoreContext := getScoreContext(trait, score, language)
	config := getLanguageConfig(language)

	return buildPromptForLanguage(traitName, score, scoreDescription, scoreContext, config)
}

// getTraitName returns the trait name in the specified language
func getTraitName(trait domain.Trait, language string) string {
	traitNames := map[string]map[domain.Trait]string{
		"de": {
			domain.TraitExtraversion:       "Extraversion",
			domain.TraitAgreeableness:      "Verträglichkeit",
			domain.TraitConscientiousness:  "Gewissenhaftigkeit",
			domain.TraitEmotionalStability: "Emotionale Stabilität",
			domain.TraitOpenness:           "Offenheit",
		},
		"en": {
			domain.TraitExtraversion:       "Extraversion",
			domain.TraitAgreeableness:      "Agreeableness",
			domain.TraitConscientiousness:  "Conscientiousness",
			domain.TraitEmotionalStability: "Emotional Stability",
			domain.TraitOpenness:           "Openness",
		},
		"tr": {
			domain.TraitExtraversion:       "Dışa Dönüklük",
			domain.TraitAgreeableness:      "Uyumluluk",
			domain.TraitConscientiousness:  "Sorumluluk",
			domain.TraitEmotionalStability: "Duygusal Denge",
			domain.TraitOpenness:           "Deneyime Açıklık",
		},
		"ar": {
			domain.TraitExtraversion:       "الانبساطية",
			domain.TraitAgreeableness:      "الوفاقية",
			domain.TraitConscientiousness:  "الضمير الحي",
			domain.TraitEmotionalStability: "الاستقرار العاطفي",
			domain.TraitOpenness:           "الانفتاح",
		},
		"ru": {
			domain.TraitExtraversion:       "Экстраверсия",
			domain.TraitAgreeableness:      "Доброжелательность",
			domain.TraitConscientiousness:  "Добросовестность",
			domain.TraitEmotionalStability: "Эмоциональная стабильность",
			domain.TraitOpenness:           "Открытость опыту",
		},
		"pl": {
			domain.TraitExtraversion:       "Ekstrawersja",
			domain.TraitAgreeableness:      "Ugodowość",
			domain.TraitConscientiousness:  "Sumienność",
			domain.TraitEmotionalStability: "Stabilność emocjonalna",
			domain.TraitOpenness:           "Otwartość na doświadczenia",
		},
		"ro": {
			domain.TraitExtraversion:       "Extraversie",
			domain.TraitAgreeableness:      "Agreabilitate",
			domain.TraitConscientiousness:  "Conștiinciozitate",
			domain.TraitEmotionalStability: "Stabilitate emoțională",
			domain.TraitOpenness:           "Deschidere",
		},
		"it": {
			domain.TraitExtraversion:       "Estroversione",
			domain.TraitAgreeableness:      "Amicalità",
			domain.TraitConscientiousness:  "Coscienziosità",
			domain.TraitEmotionalStability: "Stabilità emotiva",
			domain.TraitOpenness:           "Apertura mentale",
		},
		"uk": {
			domain.TraitExtraversion:       "Екстраверсія",
			domain.TraitAgreeableness:      "Доброзичливість",
			domain.TraitConscientiousness:  "Сумлінність",
			domain.TraitEmotionalStability: "Емоційна стабільність",
			domain.TraitOpenness:           "Відкритість досвіду",
		},
		"bg": {
			domain.TraitExtraversion:       "Екстраверсия",
			domain.TraitAgreeableness:      "Сговорчивост",
			domain.TraitConscientiousness:  "Съвестност",
			domain.TraitEmotionalStability: "Емоционална стабилност",
			domain.TraitOpenness:           "Отвореност",
		},
	}

	if names, ok := traitNames[language]; ok {
		if name, ok := names[trait]; ok {
			return name
		}
	}
	// Default to German
	return traitNames["de"][trait]
}

// buildPromptForLanguage generates the interpretation prompt in the specified language
func buildPromptForLanguage(traitName string, score float64, scoreDescription, scoreContext string, config LanguageConfig) string {
	switch config.ResponseLanguage {
	case "English":
		return fmt.Sprintf(`Write a personal, encouraging interpretation for the trait "%s" 
with a score of %.0f out of 100.

For context: A score of %.0f means %s.

%s

Write this as a flowing, connected text in exactly five sections.
Each section should have 3-5 sentences that flow naturally into the next.
Do NOT use bullet points, lists, or numbering.
Use casual, relatable language – like you're talking to a friend, not writing a textbook.

Structure your text with these five headings:

## What this means for you

Help them recognize themselves in this trait. Describe how it might show up in their daily life – 
at school, with friends, at home, during hobbies, or even just scrolling their phone. 
Keep it neutral and judgment-free. They should read this and think "yeah, that sounds like me."

## What you're probably good at

Highlight the strengths and skills that come with this trait.
Think about things like being a good listener, thinking outside the box, staying focused, 
keeping calm, hyping others up, or planning things out. 
Don't mention specific careers – just focus on skills that are useful pretty much anywhere.
Even lower scores come with real strengths that are worth noticing.

## What can be tricky sometimes

Be real about the situations where this trait might make things harder.
Use words like "can", "sometimes", or "in some situations" – no doom and gloom.
This isn't about problems, it's about understanding that every trait has trade-offs.
Help them see both sides without making them feel bad about themselves.

## Where this trait really works for you

Open up possibilities for environments and activities where this trait is actually an advantage.
Describe situations, types of work, and settings where this trait really shines – 
like working with people, creative projects, organized tasks, solo work, or team stuff.
Don't name specific jobs – keep it about vibes and environments.

## Things to think about

Ask two or three questions that get them thinking about themselves.
Make these questions spark curiosity – like when do they notice this trait most? 
When does it help them out? What situations might they want to try to learn more about themselves?
Keep the questions open and inviting, not like a homework assignment.

Respond exclusively in English.`, traitName, score, score, scoreDescription, scoreContext)

	case "Türkçe":
		return fmt.Sprintf(`"%s" özelliği için kişisel, cesaretlendirici bir yorum yaz. 
100 üzerinden %.0f puan.

Bağlam için: %.0f puan %s anlamına gelir.

%s

Bunu tam olarak beş bölümde akıcı, bağlantılı bir metin olarak yaz.
Her bölüm doğal bir şekilde bir sonrakine geçen 3-5 cümle içermeli.
Madde işaretleri, listeler veya numaralandırma KULLANMA.
Samimi, günlük bir dil kullan – bir arkadaşla konuşuyormuş gibi, ders kitabı yazmıyormuş gibi.

Metninizi şu beş başlıkla yapılandır:

## Bu senin için ne anlama geliyor

Bu özellikte kendilerini tanımalarına yardımcı ol. Günlük hayatlarında nasıl ortaya çıkabileceğini anlat – 
okulda, arkadaşlarla, evde, hobiler sırasında veya sadece telefonda gezinirken. 
Tarafsız ve yargısız tut. Bunu okuyup "evet, bu bana benziyor" demeli.

## Muhtemelen iyi olduğun şeyler

Bu özellikle gelen güçlü yönleri ve becerileri vurgula.
İyi bir dinleyici olmak, kalıpların dışında düşünmek, odaklanmak, 
sakin kalmak, başkalarını motive etmek veya işleri planlamak gibi şeyleri düşün. 
Belirli kariyerlerden bahsetme – sadece hemen hemen her yerde işe yarayan becerilere odaklan.
Düşük puanlar bile fark edilmeye değer gerçek güçlü yönlerle gelir.

## Bazen zor olabilecek şeyler

Bu özelliğin işleri zorlaştırabileceği durumlar hakkında dürüst ol.
"Olabilir", "bazen" veya "bazı durumlarda" gibi kelimeler kullan – felaket senaryosu yok.
Bu sorunlarla ilgili değil, her özelliğin artıları ve eksileri olduğunu anlamakla ilgili.
Kendileri hakkında kötü hissetmeden her iki tarafı görmelerine yardımcı ol.

## Bu özelliğin gerçekten işe yaradığı yerler

Bu özelliğin aslında bir avantaj olduğu ortamlar ve aktiviteler için olasılıklar aç.
Bu özelliğin gerçekten parladığı durumları, çalışma türlerini ve ortamları anlat – 
insanlarla çalışmak, yaratıcı projeler, organize görevler, solo çalışma veya takım işleri gibi.
Belirli işleri söyleme – ortamlar ve atmosfer hakkında tut.

## Düşünülecek şeyler

Kendileri hakkında düşünmelerini sağlayacak iki veya üç soru sor.
Bu soruların merak uyandırmasını sağla – bu özelliği en çok ne zaman fark ediyorlar? 
Ne zaman işlerine yarıyor? Kendileri hakkında daha fazla şey öğrenmek için hangi durumları denemek isteyebilirler?
Soruları açık ve davetkar tut, ödev gibi değil.

Yanıtı yalnızca Türkçe olarak ver.`, traitName, score, score, scoreDescription, scoreContext)

	case "العربية":
		return fmt.Sprintf(`اكتب تفسيرًا شخصيًا ومشجعًا لسمة "%s" 
بدرجة %.0f من 100.

للتوضيح: درجة %.0f تعني %s.

%s

اكتب هذا كنص متدفق ومترابط في خمسة أقسام بالضبط.
كل قسم يجب أن يحتوي على 3-5 جمل تتدفق بشكل طبيعي إلى القسم التالي.
لا تستخدم النقاط أو القوائم أو الترقيم.
استخدم لغة عادية وقريبة – كأنك تتحدث مع صديق، وليس ككتاب دراسي.

نظم نصك بهذه العناوين الخمسة:

## ماذا يعني هذا لك

ساعدهم على التعرف على أنفسهم في هذه السمة. صف كيف قد تظهر في حياتهم اليومية – 
في المدرسة، مع الأصدقاء، في البيت، أثناء الهوايات، أو حتى أثناء تصفح الهاتف. 
اجعلها محايدة وخالية من الأحكام. يجب أن يقرأوا هذا ويفكروا "نعم، هذا يشبهني".

## ما أنت ربما جيد فيه

أبرز نقاط القوة والمهارات التي تأتي مع هذه السمة.
فكر في أشياء مثل أن تكون مستمعًا جيدًا، التفكير خارج الصندوق، البقاء مركزًا، 
الحفاظ على الهدوء، تشجيع الآخرين، أو التخطيط للأمور. 
لا تذكر مهنًا محددة – فقط ركز على المهارات المفيدة في كل مكان تقريبًا.
حتى الدرجات المنخفضة تأتي مع نقاط قوة حقيقية تستحق الملاحظة.

## ما يمكن أن يكون صعبًا أحيانًا

كن صادقًا بشأن المواقف التي قد تجعل فيها هذه السمة الأمور أصعب.
استخدم كلمات مثل "يمكن"، "أحيانًا"، أو "في بعض المواقف" – بدون تشاؤم.
هذا ليس عن المشاكل، إنه عن فهم أن كل سمة لها مزايا وعيوب.
ساعدهم على رؤية الجانبين دون أن يشعروا بالسوء تجاه أنفسهم.

## أين تعمل هذه السمة لصالحك حقًا

افتح إمكانيات للبيئات والأنشطة حيث تكون هذه السمة ميزة فعلية.
صف المواقف وأنواع العمل والإعدادات حيث تتألق هذه السمة حقًا – 
مثل العمل مع الناس، المشاريع الإبداعية، المهام المنظمة، العمل الفردي، أو العمل الجماعي.
لا تذكر وظائف محددة – اجعلها عن الأجواء والبيئات.

## أشياء للتفكير فيها

اطرح سؤالين أو ثلاثة أسئلة تجعلهم يفكرون في أنفسهم.
اجعل هذه الأسئلة تثير الفضول – مثل متى يلاحظون هذه السمة أكثر؟ 
متى تساعدهم؟ ما المواقف التي قد يرغبون في تجربتها لمعرفة المزيد عن أنفسهم؟
اجعل الأسئلة مفتوحة وودية، وليست كواجب منزلي.

أجب باللغة العربية حصريًا.`, traitName, score, score, scoreDescription, scoreContext)

	case "русский":
		return fmt.Sprintf(`Напиши личную, ободряющую интерпретацию для черты "%s" 
с показателем %.0f из 100.

Для контекста: показатель %.0f означает %s.

%s

Напиши это как плавный, связный текст ровно в пяти разделах.
Каждый раздел должен содержать 3-5 предложений, которые естественно переходят в следующий.
НЕ используй маркированные списки, перечисления или нумерацию.
Используй простой, понятный язык — как будто разговариваешь с другом, а не пишешь учебник.

Структурируй текст с этими пятью заголовками:

## Что это значит для тебя

Помоги им узнать себя в этой черте. Опиши, как она может проявляться в их повседневной жизни — 
в школе, с друзьями, дома, во время хобби или даже просто листая ленту в телефоне. 
Пиши нейтрально и без осуждения. Они должны прочитать это и подумать «да, это про меня».

## В чём ты, скорее всего, хорош

Подчеркни сильные стороны и навыки, которые идут с этой чертой.
Подумай о таких вещах как умение слушать, нестандартное мышление, способность концентрироваться, 
сохранять спокойствие, поддерживать других или планировать дела. 
Не упоминай конкретные профессии — просто фокусируйся на навыках, которые полезны практически везде.
Даже низкие показатели несут реальные сильные стороны, которые стоит заметить.

## Что иногда может быть непросто

Будь честен насчёт ситуаций, где эта черта может усложнять жизнь.
Используй слова вроде «может», «иногда» или «в некоторых ситуациях» — без мрачных прогнозов.
Это не про проблемы, а про понимание, что у каждой черты есть свои плюсы и минусы.
Помоги им увидеть обе стороны, не заставляя чувствовать себя плохо.

## Где эта черта реально работает на тебя

Открой возможности для обстановки и занятий, где эта черта — реальное преимущество.
Опиши ситуации, виды деятельности и среды, где эта черта по-настоящему раскрывается — 
например, работа с людьми, творческие проекты, структурированные задачи, работа в одиночку или в команде.
Не называй конкретные профессии — говори об атмосфере и обстановке.

## Над чем подумать

Задай два-три вопроса, которые заставят их задуматься о себе.
Пусть эти вопросы пробуждают любопытство — например, когда они замечают эту черту больше всего? 
Когда она им помогает? Какие ситуации они хотели бы попробовать, чтобы узнать себя лучше?
Пусть вопросы будут открытыми и располагающими, не как домашнее задание.

Отвечай исключительно на русском языке.`, traitName, score, score, scoreDescription, scoreContext)

	case "polski":
		return fmt.Sprintf(`Napisz osobistą, zachęcającą interpretację dla cechy "%s" 
z wynikiem %.0f na 100.

Dla kontekstu: wynik %.0f oznacza %s.

%s

Napisz to jako płynny, powiązany tekst w dokładnie pięciu sekcjach.
Każda sekcja powinna zawierać 3-5 zdań, które naturalnie przechodzą w następną.
NIE używaj punktów, list ani numeracji.
Używaj swobodnego, codziennego języka – jakbyś rozmawiał z kumplem, nie pisał podręcznik.

Uporządkuj tekst za pomocą tych pięciu nagłówków:

## Co to dla ciebie oznacza

Pomóż im rozpoznać siebie w tej cesze. Opisz, jak może się ona przejawiać w ich codziennym życiu – 
w szkole, z przyjaciółmi, w domu, podczas hobby, a nawet podczas scrollowania telefonu. 
Pisz neutralnie i bez oceniania. Powinni to przeczytać i pomyśleć "tak, to brzmi jak ja".

## W czym prawdopodobnie jesteś dobry

Podkreśl mocne strony i umiejętności, które idą z tą cechą.
Pomyśl o rzeczach takich jak bycie dobrym słuchaczem, myślenie nieszablonowe, utrzymywanie skupienia, 
zachowanie spokoju, motywowanie innych lub planowanie. 
Nie wymieniaj konkretnych karier – po prostu skup się na umiejętnościach przydatnych praktycznie wszędzie.
Nawet niższe wyniki niosą ze sobą prawdziwe mocne strony warte zauważenia.

## Co czasami może być trudne

Bądź szczery co do sytuacji, w których ta cecha może utrudniać życie.
Używaj słów jak "może", "czasami" lub "w niektórych sytuacjach" – bez czarnowidztwa.
To nie jest o problemach, to o zrozumieniu, że każda cecha ma swoje plusy i minusy.
Pomóż im zobaczyć obie strony bez poczucia, że coś z nimi nie tak.

## Gdzie ta cecha naprawdę działa na twoją korzyść

Otwórz możliwości dla środowisk i aktywności, gdzie ta cecha jest faktycznie atutem.
Opisz sytuacje, rodzaje pracy i otoczenia, w których ta cecha naprawdę błyszczy – 
jak praca z ludźmi, kreatywne projekty, uporządkowane zadania, praca solo lub w zespole.
Nie wymieniaj konkretnych zawodów – mów o klimacie i środowiskach.

## Rzeczy do przemyślenia

Zadaj dwa lub trzy pytania, które skłonią ich do myślenia o sobie.
Niech te pytania budzą ciekawość – kiedy zauważają tę cechę najbardziej? 
Kiedy im pomaga? Jakie sytuacje mogliby chcieć wypróbować, żeby dowiedzieć się więcej o sobie?
Niech pytania będą otwarte i zachęcające, nie jak zadanie domowe.

Odpowiedz wyłącznie po polsku.`, traitName, score, score, scoreDescription, scoreContext)

	case "română":
		return fmt.Sprintf(`Scrie o interpretare personală, încurajatoare pentru trăsătura "%s" 
cu un scor de %.0f din 100.

Pentru context: un scor de %.0f înseamnă %s.

%s

Scrie asta ca un text fluent, conectat în exact cinci secțiuni.
Fiecare secțiune ar trebui să aibă 3-5 propoziții care curg natural în următoarea.
NU folosi puncte, liste sau numerotare.
Folosește un limbaj relaxat, apropiat – ca și cum ai vorbi cu un prieten, nu scrii un manual.

Structurează textul cu aceste cinci titluri:

## Ce înseamnă asta pentru tine

Ajută-i să se recunoască în această trăsătură. Descrie cum ar putea să apară în viața lor de zi cu zi – 
la școală, cu prietenii, acasă, în timpul hobby-urilor, sau chiar doar scrollând pe telefon. 
Păstrează-o neutră și fără judecăți. Ar trebui să citească asta și să gândească "da, asta sună ca mine".

## La ce ești probabil bun

Evidențiază punctele forte și abilitățile care vin cu această trăsătură.
Gândește-te la lucruri precum a fi un bun ascultător, a gândi outside the box, a rămâne concentrat, 
a păstra calmul, a-i motiva pe alții, sau a planifica lucruri. 
Nu menționa cariere specifice – doar concentrează-te pe abilități utile cam peste tot.
Chiar și scorurile mai mici vin cu puncte forte reale care merită observate.

## Ce poate fi uneori dificil

Fii sincer despre situațiile în care această trăsătură ar putea îngreuna lucrurile.
Folosește cuvinte precum "poate", "uneori", sau "în unele situații" – fără catastrofism.
Nu e despre probleme, e despre înțelegerea că fiecare trăsătură are avantaje și dezavantaje.
Ajută-i să vadă ambele părți fără să se simtă prost în legătură cu ei înșiși.

## Unde această trăsătură chiar funcționează pentru tine

Deschide posibilități pentru medii și activități unde această trăsătură e de fapt un avantaj.
Descrie situații, tipuri de muncă și setări unde această trăsătură chiar strălucește – 
precum lucrul cu oamenii, proiecte creative, sarcini organizate, muncă solo, sau chestii de echipă.
Nu numi joburi specifice – păstrează-o despre vibe-uri și medii.

## Lucruri la care să te gândești

Pune două sau trei întrebări care să-i facă să se gândească la ei înșiși.
Fă aceste întrebări să stârnească curiozitatea – când observă această trăsătură cel mai mult? 
Când îi ajută? Ce situații ar putea vrea să încerce pentru a afla mai multe despre ei înșiși?
Păstrează întrebările deschise și primitoare, nu ca un tema pentru acasă.

Răspunde exclusiv în română.`, traitName, score, score, scoreDescription, scoreContext)

	case "italiano":
		return fmt.Sprintf(`Scrivi un'interpretazione personale e incoraggiante per il tratto "%s" 
con un punteggio di %.0f su 100.

Per contesto: un punteggio di %.0f significa %s.

%s

Scrivi questo come un testo fluido e connesso in esattamente cinque sezioni.
Ogni sezione dovrebbe avere 3-5 frasi che scorrono naturalmente nella successiva.
NON usare punti elenco, liste o numerazione.
Usa un linguaggio rilassato e accessibile – come se parlassi con un amico, non scrivessi un manuale.

Struttura il tuo testo con questi cinque titoli:

## Cosa significa questo per te

Aiutali a riconoscersi in questo tratto. Descrivi come potrebbe manifestarsi nella loro vita quotidiana – 
a scuola, con gli amici, a casa, durante gli hobby, o anche solo scrollando il telefono. 
Mantienilo neutrale e senza giudizi. Dovrebbero leggere questo e pensare "sì, mi somiglia".

## In cosa sei probabilmente bravo

Evidenzia i punti di forza e le competenze che accompagnano questo tratto.
Pensa a cose come essere un buon ascoltatore, pensare fuori dagli schemi, restare concentrato, 
mantenere la calma, motivare gli altri, o pianificare le cose. 
Non menzionare carriere specifiche – concentrati solo sulle competenze utili praticamente ovunque.
Anche i punteggi più bassi portano con sé veri punti di forza che vale la pena notare.

## Cosa può essere complicato a volte

Sii onesto riguardo alle situazioni in cui questo tratto potrebbe rendere le cose più difficili.
Usa parole come "può", "a volte", o "in alcune situazioni" – niente catastrofismo.
Non si tratta di problemi, ma di capire che ogni tratto ha i suoi pro e contro.
Aiutali a vedere entrambi i lati senza farli sentire male con se stessi.

## Dove questo tratto funziona davvero per te

Apri possibilità per ambienti e attività dove questo tratto è effettivamente un vantaggio.
Descrivi situazioni, tipi di lavoro e contesti dove questo tratto brilla davvero – 
come lavorare con le persone, progetti creativi, compiti organizzati, lavoro in solitaria, o in team.
Non nominare lavori specifici – parla di atmosfere e ambienti.

## Cose su cui riflettere

Fai due o tre domande che li facciano pensare a se stessi.
Fai in modo che queste domande suscitino curiosità – quando notano di più questo tratto? 
Quando li aiuta? Quali situazioni potrebbero voler provare per scoprire di più su se stessi?
Mantieni le domande aperte e invitanti, non come un compito a casa.

Rispondi esclusivamente in italiano.`, traitName, score, score, scoreDescription, scoreContext)

	case "українська":
		return fmt.Sprintf(`Напиши особисту, підбадьорливу інтерпретацію для риси "%s" 
з показником %.0f зі 100.

Для контексту: показник %.0f означає %s.

%s

Напиши це як плавний, зв'язний текст рівно у п'яти розділах.
Кожен розділ має містити 3-5 речень, які природно переходять в наступний.
НЕ використовуй марковані списки, переліки чи нумерацію.
Використовуй просту, зрозумілу мову — наче розмовляєш з другом, а не пишеш підручник.

Структуруй текст з цими п'ятьма заголовками:

## Що це означає для тебе

Допоможи їм впізнати себе в цій рисі. Опиши, як вона може проявлятися в їхньому повсякденному житті — 
у школі, з друзями, вдома, під час хобі чи навіть просто гортаючи стрічку в телефоні. 
Пиши нейтрально і без засуджень. Вони повинні прочитати це і подумати «так, це про мене».

## В чому ти, напевно, хороший

Підкресли сильні сторони та навички, які йдуть з цією рисою.
Подумай про такі речі як вміння слухати, нестандартне мислення, здатність концентруватися, 
зберігати спокій, підтримувати інших чи планувати справи. 
Не згадуй конкретні професії — просто фокусуйся на навичках, які корисні практично всюди.
Навіть низькі показники несуть реальні сильні сторони, які варто помітити.

## Що іноді може бути непросто

Будь чесним щодо ситуацій, де ця риса може ускладнювати життя.
Використовуй слова на кшталт «може», «іноді» чи «в деяких ситуаціях» — без похмурих прогнозів.
Це не про проблеми, а про розуміння, що в кожної риси є свої плюси і мінуси.
Допоможи їм побачити обидві сторони, не змушуючи почуватися погано.

## Де ця риса реально працює на тебе

Відкрий можливості для обстановки та занять, де ця риса — реальна перевага.
Опиши ситуації, види діяльності та середовища, де ця риса по-справжньому розкривається — 
наприклад, робота з людьми, творчі проекти, структуровані завдання, робота на самоті чи в команді.
Не називай конкретні професії — говори про атмосферу та обстановку.

## Над чим подумати

Постав два-три питання, які змусять їх задуматися про себе.
Нехай ці питання пробуджують цікавість — наприклад, коли вони помічають цю рису найбільше? 
Коли вона їм допомагає? Які ситуації вони хотіли б спробувати, щоб дізнатися себе краще?
Нехай питання будуть відкритими та привітними, не як домашнє завдання.

Відповідай виключно українською мовою.`, traitName, score, score, scoreDescription, scoreContext)

	case "български":
		return fmt.Sprintf(`Напиши лична, насърчаваща интерпретация за чертата "%s" 
с резултат %.0f от 100.

За контекст: резултат %.0f означава %s.

%s

Напиши това като плавен, свързан текст в точно пет раздела.
Всеки раздел трябва да има 3-5 изречения, които преминават естествено в следващия.
НЕ използвай точки, списъци или номерация.
Използвай спокоен, ежедневен език – като че говориш с приятел, а не пишеш учебник.

Структурирай текста си с тези пет заглавия:

## Какво означава това за теб

Помогни им да се разпознаят в тази черта. Опиши как може да се проявява в ежедневието им – 
в училище, с приятели, вкъщи, по време на хобита, или дори просто скролвайки в телефона. 
Дръж го неутрално и без осъждане. Трябва да прочетат това и да помислят "да, това прилича на мен".

## В какво вероятно си добър

Подчертай силните страни и уменията, които идват с тази черта.
Помисли за неща като да си добър слушател, да мислиш нестандартно, да оставаш фокусиран, 
да запазваш спокойствие, да мотивираш другите или да планираш нещата. 
Не споменавай конкретни кариери – просто се фокусирай върху умения, полезни почти навсякъде.
Дори по-ниските резултати идват с истински силни страни, които си струва да забележиш.

## Какво понякога може да е трудно

Бъди честен за ситуациите, в които тази черта може да затруднява нещата.
Използвай думи като "може", "понякога" или "в някои ситуации" – без мрачни прогнози.
Това не е за проблеми, а за разбиране, че всяка черта има своите плюсове и минуси.
Помогни им да видят и двете страни, без да се чувстват зле за себе си.

## Къде тази черта наистина работи за теб

Отвори възможности за среди и дейности, където тази черта е истинско предимство.
Опиши ситуации, видове работа и обстановки, където тази черта наистина блести – 
като работа с хора, творчески проекти, организирани задачи, самостоятелна работа или екипни неща.
Не назовавай конкретни професии – дръж се за атмосферата и средите.

## Неща за размисъл

Задай два-три въпроса, които ще ги накарат да мислят за себе си.
Направи тези въпроси да събуждат любопитство – кога забелязват тази черта най-много? 
Кога им помага? Какви ситуации биха искали да опитат, за да научат повече за себе си?
Дръж въпросите отворени и приветливи, не като домашна работа.

Отговори изключително на български.`, traitName, score, score, scoreDescription, scoreContext)

	default: // German (default)
		return fmt.Sprintf(`Schreibe eine persönliche, ermutigende Interpretation für die Eigenschaft "%s" 
mit einem Wert von %.0f von 100.

Zur Einordnung: Ein Wert von %.0f bedeutet %s.

%s

Schreibe das als fließenden, zusammenhängenden Text in genau fünf Abschnitten.
Jeder Abschnitt soll 3-5 Sätze umfassen und natürlich in den nächsten übergehen.
Verwende KEINE Stichpunkte, Aufzählungen oder Nummerierungen.
Nutze eine lockere, alltägliche Sprache – als würdest du mit einem Freund reden, nicht ein Lehrbuch schreiben.

Strukturiere deinen Text mit diesen fünf Überschriften:

## Was das für dich bedeutet

Hilf ihnen, sich in dieser Eigenschaft wiederzuerkennen. Beschreibe, wie sie sich im Alltag zeigen könnte – 
in der Schule, mit Freunden, zuhause, bei Hobbys oder auch einfach beim Scrollen am Handy. 
Halte es neutral und wertfrei. Sie sollen das lesen und denken "ja, das klingt nach mir".

## Das kannst du wahrscheinlich gut

Hebe die Stärken und Fähigkeiten hervor, die mit dieser Eigenschaft kommen.
Denk an Dinge wie gut zuhören können, kreativ denken, fokussiert bleiben, 
ruhig bleiben, andere motivieren oder Dinge planen. 
Nenne keine konkreten Berufe – fokussiere dich auf Fähigkeiten, die praktisch überall nützlich sind.
Auch niedrigere Werte bringen echte Stärken mit sich, die es wert sind, bemerkt zu werden.

## Das kann manchmal knifflig sein

Sei ehrlich darüber, in welchen Situationen diese Eigenschaft es schwieriger machen kann.
Nutze Wörter wie "kann", "manchmal" oder "in manchen Situationen" – kein Schwarzmalen.
Es geht nicht um Probleme, sondern darum zu verstehen, dass jede Eigenschaft ihre Vor- und Nachteile hat.
Hilf ihnen, beide Seiten zu sehen, ohne dass sie sich schlecht fühlen.

## Wo diese Eigenschaft richtig gut für dich funktioniert

Öffne Möglichkeiten für Umgebungen und Aktivitäten, wo diese Eigenschaft ein echter Vorteil ist.
Beschreibe Situationen, Arbeitsarten und Settings, wo diese Eigenschaft richtig aufblüht – 
wie Arbeit mit Menschen, kreative Projekte, strukturierte Aufgaben, Arbeit allein oder im Team.
Nenne keine konkreten Jobs – bleib bei Vibes und Umgebungen.

## Zum Nachdenken

Stelle zwei bis drei Fragen, die sie über sich selbst nachdenken lassen.
Diese Fragen sollen neugierig machen – wann bemerken sie diese Eigenschaft am meisten? 
Wann hilft sie ihnen? Welche Situationen würden sie gerne ausprobieren, um mehr über sich zu erfahren?
Halte die Fragen offen und einladend, nicht wie eine Hausaufgabe.

Antworte ausschließlich auf Deutsch.`, traitName, score, score, scoreDescription, scoreContext)
	}
}

// describeScore returns a qualitative description of the score level in the specified language
func describeScore(score float64, language string) string {
	descriptions := map[string]map[string]string{
		"de": {
			"very_high": "diese Eigenschaft ist bei dir richtig stark ausgeprägt",
			"high":      "diese Eigenschaft zeigt sich bei dir ziemlich deutlich",
			"medium":    "du bist bei dieser Eigenschaft irgendwo in der Mitte – ziemlich ausgewogen",
			"low":       "diese Eigenschaft ist bei dir eher dezent vorhanden",
			"very_low":  "diese Eigenschaft ist bei dir nicht so stark ausgeprägt – und das ist völlig okay",
		},
		"en": {
			"very_high": "this trait is really strong in you",
			"high":      "this trait shows up quite a bit in how you are",
			"medium":    "you're somewhere in the middle on this one – pretty balanced",
			"low":       "this trait is more on the subtle side for you",
			"very_low":  "this trait isn't super prominent for you – and that's totally fine",
		},
		"tr": {
			"very_high": "bu özellik sende gerçekten güçlü",
			"high":      "bu özellik sende oldukça belirgin",
			"medium":    "bu özellikte ortada bir yerdesin – oldukça dengeli",
			"low":       "bu özellik sende biraz daha arka planda",
			"very_low":  "bu özellik sende çok baskın değil – ve bu tamamen normal",
		},
		"ar": {
			"very_high": "هذه السمة قوية جداً فيك",
			"high":      "هذه السمة واضحة جداً في شخصيتك",
			"medium":    "أنت في المنتصف في هذه السمة – متوازن تماماً",
			"low":       "هذه السمة أقل وضوحاً فيك",
			"very_low":  "هذه السمة ليست بارزة جداً فيك – وهذا طبيعي تماماً",
		},
		"ru": {
			"very_high": "эта черта у тебя реально сильная",
			"high":      "эта черта у тебя довольно заметная",
			"medium":    "ты где-то посередине в этой черте – довольно сбалансированно",
			"low":       "эта черта у тебя скорее в тени",
			"very_low":  "эта черта у тебя не особо выражена – и это абсолютно нормально",
		},
		"pl": {
			"very_high": "ta cecha jest u ciebie naprawdę silna",
			"high":      "ta cecha jest u ciebie całkiem wyraźna",
			"medium":    "jesteś gdzieś pośrodku w tej cesze – całkiem zbalansowany",
			"low":       "ta cecha jest u ciebie raczej subtelna",
			"very_low":  "ta cecha nie jest u ciebie super widoczna – i to jest w porządku",
		},
		"ro": {
			"very_high": "această trăsătură e foarte puternică la tine",
			"high":      "această trăsătură se vede destul de clar la tine",
			"medium":    "ești undeva la mijloc cu asta – destul de echilibrat",
			"low":       "această trăsătură e mai subtilă la tine",
			"very_low":  "această trăsătură nu e super evidentă la tine – și asta e perfect ok",
		},
		"it": {
			"very_high": "questo tratto è davvero forte in te",
			"high":      "questo tratto si vede abbastanza in come sei",
			"medium":    "sei più o meno nel mezzo su questo – abbastanza equilibrato",
			"low":       "questo tratto è più sottile in te",
			"very_low":  "questo tratto non è super evidente in te – e va benissimo così",
		},
		"uk": {
			"very_high": "ця риса у тебе реально сильна",
			"high":      "ця риса у тебе доволі помітна",
			"medium":    "ти десь посередині в цій рисі – доволі збалансовано",
			"low":       "ця риса у тебе скоріше в тіні",
			"very_low":  "ця риса у тебе не особливо виражена – і це абсолютно нормально",
		},
		"bg": {
			"very_high": "тази черта е наистина силна при теб",
			"high":      "тази черта се проявява доста ясно при теб",
			"medium":    "ти си някъде по средата при тази черта – доста балансирано",
			"low":       "тази черта е по-скоро дискретна при теб",
			"very_low":  "тази черта не е супер изразена при теб – и това е напълно нормално",
		},
	}

	// Get the description map for the language, defaulting to German
	langDescriptions, ok := descriptions[language]
	if !ok {
		langDescriptions = descriptions["de"]
	}

	var level string
	switch {
	case score >= 80:
		level = "very_high"
	case score >= 60:
		level = "high"
	case score >= 40:
		level = "medium"
	case score >= 20:
		level = "low"
	default:
		level = "very_low"
	}

	return langDescriptions[level]
}

// getScoreContext provides trait-specific context to help the AI understand
// what high/low scores mean for each dimension.
func getScoreContext(trait domain.Trait, score float64, language string) string {
	isHigh := score >= 60
	isLow := score < 40

	// Get context based on language
	contexts := getTraitContexts(language)

	var level string
	if isHigh {
		level = "high"
	} else if isLow {
		level = "low"
	} else {
		level = "medium"
	}

	if traitContexts, ok := contexts[trait]; ok {
		if context, ok := traitContexts[level]; ok {
			return context
		}
	}

	return ""
}

// getTraitContexts returns trait-specific context strings for a given language
func getTraitContexts(language string) map[domain.Trait]map[string]string {
	switch language {
	case "en":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `For Extraversion, a high score means you get energy from being around people, enjoy being in the spotlight, and feel at home in groups.`,
				"low":    `For Extraversion, a low score means you recharge during quiet time alone and prefer real talk with a few close people over big social scenes.`,
				"medium": `For Extraversion, a medium score means you can switch between hanging out with people and enjoying your own company – whatever fits the moment.`,
			},
			domain.TraitAgreeableness: {
				"high":   `For Agreeableness, a high score means you care about keeping the peace, like working with others, and often think about what other people need.`,
				"low":    `For Agreeableness, a low score means you speak your mind, don't back down from disagreements, and don't let others easily sway you.`,
				"medium": `For Agreeableness, a medium score means you can be a team player when it matters and stand up for yourself when you need to.`,
			},
			domain.TraitConscientiousness: {
				"high":   `For Conscientiousness, a high score means you like to stay organized, get stuff done on time, and stick with long-term goals.`,
				"low":    `For Conscientiousness, a low score means you're flexible and go with the flow, don't stress too much about plans, and adapt easily when things change.`,
				"medium": `For Conscientiousness, a medium score means you can be organized when it counts but also roll with the punches when needed.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `For Emotional Stability, a high score means you stay pretty chill under pressure, don't get too thrown off by setbacks, and keep your cool.`,
				"low":    `For Emotional Stability, a low score means you feel things deeply, pick up on vibes around you, and have strong emotional awareness.`,
				"medium": `For Emotional Stability, a medium score means you can feel things intensely but also know how to manage those feelings.`,
			},
			domain.TraitOpenness: {
				"high":   `For Openness, a high score means you're curious about trying new things, love thinking creatively, and get excited by big ideas.`,
				"low":    `For Openness, a low score means you're practical and down-to-earth, appreciate what's tried and true, and like having familiar routines.`,
				"medium": `For Openness, a medium score means you enjoy new experiences but also know when to stick with what works.`,
			},
		}

	case "tr":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Dışa Dönüklük için yüksek puan, insanların arasında enerji topladığın, ilgi odağı olmaktan keyif aldığın ve gruplarda kendini evinde hissettiğin anlamına gelir.`,
				"low":    `Dışa Dönüklük için düşük puan, yalnız geçirdiğin sakin zamanlarda şarj olduğun ve büyük sosyal ortamlar yerine birkaç yakın arkadaşla gerçek sohbetleri tercih ettiğin anlamına gelir.`,
				"medium": `Dışa Dönüklük için orta puan, insanlarla takılmak ve kendi halinde olmak arasında geçiş yapabildiğin anlamına gelir – ana ne uygunsa.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Uyumluluk için yüksek puan, barışı korumayı önemsediğin, başkalarıyla çalışmayı sevdiğin ve genellikle diğer insanların neye ihtiyacı olduğunu düşündüğün anlamına gelir.`,
				"low":    `Uyumluluk için düşük puan, fikrini söylediğin, anlaşmazlıklardan geri adım atmadığın ve başkalarının seni kolayca etkilemesine izin vermediğin anlamına gelir.`,
				"medium": `Uyumluluk için orta puan, önemli olduğunda takım oyuncusu olabildiğin ve gerektiğinde kendini savunabildiğin anlamına gelir.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Sorumluluk için yüksek puan, organize kalmayı, işleri zamanında bitirmeyi ve uzun vadeli hedeflere bağlı kalmayı sevdiğin anlamına gelir.`,
				"low":    `Sorumluluk için düşük puan, esnek olduğun ve akışa bıraktığın, planlar konusunda çok stres yapmadığın ve işler değiştiğinde kolayca adapte olduğun anlamına gelir.`,
				"medium": `Sorumluluk için orta puan, önemli olduğunda organize olabildiğin ama gerektiğinde duruma da uyum sağlayabildiğin anlamına gelir.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Duygusal Denge için yüksek puan, baskı altında oldukça sakin kaldığın, aksiliklerden çok etkilenmediğin ve soğukkanlılığını koruduğun anlamına gelir.`,
				"low":    `Duygusal Denge için düşük puan, duyguları derinden hissettiğin, etrafındaki havayı yakaladığın ve güçlü duygusal farkındalığa sahip olduğun anlamına gelir.`,
				"medium": `Duygusal Denge için orta puan, duyguları yoğun yaşayabildiğin ama aynı zamanda onları nasıl yöneteceğini de bildiğin anlamına gelir.`,
			},
			domain.TraitOpenness: {
				"high":   `Deneyime Açıklık için yüksek puan, yeni şeyler denemeye meraklı olduğun, yaratıcı düşünmeyi sevdiğin ve büyük fikirlerden heyecan duyduğun anlamına gelir.`,
				"low":    `Deneyime Açıklık için düşük puan, pratik ve ayakları yere basan biri olduğun, denenmiş ve doğru olanı takdir ettiğin ve tanıdık rutinleri sevdiğin anlamına gelir.`,
				"medium": `Deneyime Açıklık için orta puan, yeni deneyimlerden keyif aldığın ama aynı zamanda işe yarayan şeylere ne zaman bağlı kalacağını da bildiğin anlamına gelir.`,
			},
		}

	case "ar":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `بالنسبة للانبساطية، الدرجة العالية تعني أنك تحصل على طاقتك من التواجد حول الناس، وتستمتع بأن تكون تحت الأضواء، وتشعر بالراحة في المجموعات.`,
				"low":    `بالنسبة للانبساطية، الدرجة المنخفضة تعني أنك تستعيد طاقتك في الأوقات الهادئة بمفردك وتفضل الحديث الحقيقي مع أشخاص مقربين قليلين على التجمعات الكبيرة.`,
				"medium": `بالنسبة للانبساطية، الدرجة المتوسطة تعني أنك تستطيع التبديل بين التواجد مع الناس والاستمتاع بوقتك الخاص – حسب ما يناسب اللحظة.`,
			},
			domain.TraitAgreeableness: {
				"high":   `بالنسبة للوفاقية، الدرجة العالية تعني أنك تهتم بالحفاظ على السلام، وتحب العمل مع الآخرين، وغالباً تفكر فيما يحتاجه الآخرون.`,
				"low":    `بالنسبة للوفاقية، الدرجة المنخفضة تعني أنك تقول رأيك، ولا تتراجع عن الخلافات، ولا تدع الآخرين يؤثرون عليك بسهولة.`,
				"medium": `بالنسبة للوفاقية، الدرجة المتوسطة تعني أنك تستطيع أن تكون لاعب فريق عندما يهم الأمر وتدافع عن نفسك عندما تحتاج.`,
			},
			domain.TraitConscientiousness: {
				"high":   `بالنسبة للضمير الحي، الدرجة العالية تعني أنك تحب البقاء منظماً، وإنجاز الأمور في الوقت المحدد، والالتزام بالأهداف طويلة المدى.`,
				"low":    `بالنسبة للضمير الحي، الدرجة المنخفضة تعني أنك مرن وتتماشى مع التيار، ولا تقلق كثيراً بشأن الخطط، وتتكيف بسهولة عندما تتغير الأمور.`,
				"medium": `بالنسبة للضمير الحي، الدرجة المتوسطة تعني أنك تستطيع أن تكون منظماً عندما يهم الأمر ولكن أيضاً تتعامل مع المفاجآت عند الحاجة.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `بالنسبة للاستقرار العاطفي، الدرجة العالية تعني أنك تبقى هادئاً جداً تحت الضغط، ولا تتأثر كثيراً بالانتكاسات، وتحافظ على هدوئك.`,
				"low":    `بالنسبة للاستقرار العاطفي، الدرجة المنخفضة تعني أنك تشعر بالأشياء بعمق، وتلتقط الأجواء من حولك، ولديك وعي عاطفي قوي.`,
				"medium": `بالنسبة للاستقرار العاطفي، الدرجة المتوسطة تعني أنك تستطيع الشعور بالأشياء بشكل مكثف ولكن أيضاً تعرف كيف تدير تلك المشاعر.`,
			},
			domain.TraitOpenness: {
				"high":   `بالنسبة للانفتاح، الدرجة العالية تعني أنك فضولي لتجربة أشياء جديدة، وتحب التفكير بإبداع، وتتحمس للأفكار الكبيرة.`,
				"low":    `بالنسبة للانفتاح، الدرجة المنخفضة تعني أنك عملي وواقعي، وتقدر ما هو مجرب وصحيح، وتحب الروتين المألوف.`,
				"medium": `بالنسبة للانفتاح، الدرجة المتوسطة تعني أنك تستمتع بالتجارب الجديدة ولكن أيضاً تعرف متى تلتزم بما ينجح.`,
			},
		}

	case "ru":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Экстраверсии высокий показатель означает, что ты заряжаешься от общения с людьми, тебе нравится быть в центре внимания, и ты чувствуешь себя как рыба в воде в компаниях.`,
				"low":    `При Экстраверсии низкий показатель означает, что ты восстанавливаешься в тихие моменты наедине с собой и предпочитаешь настоящие разговоры с несколькими близкими людьми большим тусовкам.`,
				"medium": `При Экстраверсии средний показатель означает, что ты можешь переключаться между общением с людьми и своим личным временем — как подходит моменту.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Доброжелательности высокий показатель означает, что тебе важно сохранять мир, ты любишь работать с другими и часто думаешь о том, что нужно окружающим.`,
				"low":    `При Доброжелательности низкий показатель означает, что ты говоришь что думаешь, не отступаешь в спорах и не позволяешь другим легко на тебя влиять.`,
				"medium": `При Доброжелательности средний показатель означает, что ты можешь быть командным игроком, когда это важно, и отстаивать себя, когда нужно.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Добросовестности высокий показатель означает, что тебе нравится быть организованным, делать дела вовремя и держаться долгосрочных целей.`,
				"low":    `При Добросовестности низкий показатель означает, что ты гибкий и плывёшь по течению, не паришься сильно из-за планов и легко адаптируешься, когда что-то меняется.`,
				"medium": `При Добросовестности средний показатель означает, что ты можешь быть организованным, когда это важно, но и справляться с неожиданностями, когда нужно.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Эмоциональной стабильности высокий показатель означает, что ты остаёшься довольно спокойным под давлением, не сильно выбиваешься из колеи неудачами и сохраняешь хладнокровие.`,
				"low":    `При Эмоциональной стабильности низкий показатель означает, что ты глубоко чувствуешь, улавливаешь настроение вокруг и обладаешь сильным эмоциональным восприятием.`,
				"medium": `При Эмоциональной стабильности средний показатель означает, что ты можешь чувствовать интенсивно, но также знаешь, как управлять этими чувствами.`,
			},
			domain.TraitOpenness: {
				"high":   `При Открытости опыту высокий показатель означает, что тебе интересно пробовать новое, ты любишь мыслить креативно и загораешься от крутых идей.`,
				"low":    `При Открытости опыту низкий показатель означает, что ты практичный и приземлённый, ценишь проверенное временем и любишь привычные ритмы.`,
				"medium": `При Открытости опыту средний показатель означает, что тебе нравится новый опыт, но ты также знаешь, когда придерживаться того, что работает.`,
			},
		}

	case "pl":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Dla Ekstrawersji wysoki wynik oznacza, że ładujesz się energią od ludzi, lubisz być w centrum uwagi i czujesz się jak u siebie w grupach.`,
				"low":    `Dla Ekstrawersji niski wynik oznacza, że regenerujesz się w spokojne chwile sam na sam i wolisz prawdziwe rozmowy z kilkoma bliskimi osobami niż wielkie imprezy.`,
				"medium": `Dla Ekstrawersji średni wynik oznacza, że możesz przełączać się między byciem z ludźmi a cieszeniem się własnym towarzystwem – co pasuje do momentu.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Dla Ugodowości wysoki wynik oznacza, że zależy ci na utrzymaniu pokoju, lubisz pracować z innymi i często myślisz o tym, czego potrzebują inni.`,
				"low":    `Dla Ugodowości niski wynik oznacza, że mówisz co myślisz, nie wycofujesz się z nieporozumień i nie dajesz się łatwo przekonać innym.`,
				"medium": `Dla Ugodowości średni wynik oznacza, że możesz być graczem zespołowym, gdy to ważne, i stanąć w swojej obronie, gdy trzeba.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Dla Sumienności wysoki wynik oznacza, że lubisz być zorganizowany, robić rzeczy na czas i trzymać się długoterminowych celów.`,
				"low":    `Dla Sumienności niski wynik oznacza, że jesteś elastyczny i płyniesz z prądem, nie stresujesz się za bardzo planami i łatwo adaptujesz się, gdy coś się zmienia.`,
				"medium": `Dla Sumienności średni wynik oznacza, że możesz być zorganizowany, gdy to ważne, ale też radzić sobie z niespodziankami, gdy trzeba.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Dla Stabilności emocjonalnej wysoki wynik oznacza, że pozostajesz dość spokojny pod presją, nie dajesz się za bardzo wytrącić niepowodzeniami i zachowujesz zimną krew.`,
				"low":    `Dla Stabilności emocjonalnej niski wynik oznacza, że czujesz rzeczy głęboko, łapiesz nastrój wokół siebie i masz silną świadomość emocjonalną.`,
				"medium": `Dla Stabilności emocjonalnej średni wynik oznacza, że możesz czuć intensywnie, ale też wiesz, jak zarządzać tymi uczuciami.`,
			},
			domain.TraitOpenness: {
				"high":   `Dla Otwartości na doświadczenia wysoki wynik oznacza, że jesteś ciekaw próbowania nowych rzeczy, kochasz myśleć kreatywnie i ekscytujesz się wielkimi ideami.`,
				"low":    `Dla Otwartości na doświadczenia niski wynik oznacza, że jesteś praktyczny i twardo stąpasz po ziemi, cenisz to co sprawdzone i lubisz znane rutyny.`,
				"medium": `Dla Otwartości na doświadczenia średni wynik oznacza, że lubisz nowe doświadczenia, ale też wiesz, kiedy trzymać się tego, co działa.`,
			},
		}

	case "ro":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Pentru Extraversie, un scor mare înseamnă că îți iei energia din a fi cu oamenii, îți place să fii în centrul atenției și te simți ca acasă în grupuri.`,
				"low":    `Pentru Extraversie, un scor mic înseamnă că te reîncarci în timpul momentelor liniștite singur și preferi conversațiile reale cu câțiva oameni apropiați decât scenele sociale mari.`,
				"medium": `Pentru Extraversie, un scor mediu înseamnă că poți comuta între a fi cu oamenii și a te bucura de propria companie – ce se potrivește momentului.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Pentru Agreabilitate, un scor mare înseamnă că îți pasă de menținerea păcii, îți place să lucrezi cu alții și adesea te gândești la ce au nevoie alții.`,
				"low":    `Pentru Agreabilitate, un scor mic înseamnă că îți spui părerea, nu dai înapoi de la dezacorduri și nu te lași ușor influențat de alții.`,
				"medium": `Pentru Agreabilitate, un scor mediu înseamnă că poți fi un jucător de echipă când contează și să te aperi când trebuie.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Pentru Conștiinciozitate, un scor mare înseamnă că îți place să rămâi organizat, să faci lucrurile la timp și să te ții de obiectivele pe termen lung.`,
				"low":    `Pentru Conștiinciozitate, un scor mic înseamnă că ești flexibil și mergi cu fluxul, nu te stresezi prea mult cu planurile și te adaptezi ușor când lucrurile se schimbă.`,
				"medium": `Pentru Conștiinciozitate, un scor mediu înseamnă că poți fi organizat când contează dar și să te descurci cu surprizele când e nevoie.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Pentru Stabilitatea emoțională, un scor mare înseamnă că rămâi destul de chill sub presiune, nu te dai peste cap de eșecuri și îți păstrezi calmul.`,
				"low":    `Pentru Stabilitatea emoțională, un scor mic înseamnă că simți lucrurile profund, prinzi vibe-urile din jur și ai o conștientizare emoțională puternică.`,
				"medium": `Pentru Stabilitatea emoțională, un scor mediu înseamnă că poți simți lucrurile intens dar și știi cum să gestionezi acele sentimente.`,
			},
			domain.TraitOpenness: {
				"high":   `Pentru Deschidere, un scor mare înseamnă că ești curios să încerci lucruri noi, îți place să gândești creativ și te entuziasmezi de ideile mari.`,
				"low":    `Pentru Deschidere, un scor mic înseamnă că ești practic și cu picioarele pe pământ, apreciezi ce e testat și adevărat și îți plac rutinele familiare.`,
				"medium": `Pentru Deschidere, un scor mediu înseamnă că te bucuri de experiențe noi dar și știi când să te ții de ce funcționează.`,
			},
		}

	case "it":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Per l'Estroversione, un punteggio alto significa che ti carichi stando con le persone, ti piace essere sotto i riflettori e ti senti a casa nei gruppi.`,
				"low":    `Per l'Estroversione, un punteggio basso significa che ti ricarichi nei momenti tranquilli da solo e preferisci chiacchierate vere con poche persone vicine rispetto alle grandi scene sociali.`,
				"medium": `Per l'Estroversione, un punteggio medio significa che puoi passare tra stare con le persone e goderti la tua compagnia – quello che si adatta al momento.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Per l'Amicalità, un punteggio alto significa che ti importa mantenere la pace, ti piace lavorare con gli altri e spesso pensi a cosa hanno bisogno gli altri.`,
				"low":    `Per l'Amicalità, un punteggio basso significa che dici la tua, non ti tiri indietro dai disaccordi e non ti lasci influenzare facilmente dagli altri.`,
				"medium": `Per l'Amicalità, un punteggio medio significa che puoi essere un giocatore di squadra quando conta e difenderti quando serve.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Per la Coscienziosità, un punteggio alto significa che ti piace restare organizzato, fare le cose in tempo e attenerti agli obiettivi a lungo termine.`,
				"low":    `Per la Coscienziosità, un punteggio basso significa che sei flessibile e vai con il flusso, non ti stressi troppo per i piani e ti adatti facilmente quando le cose cambiano.`,
				"medium": `Per la Coscienziosità, un punteggio medio significa che puoi essere organizzato quando conta ma anche gestire gli imprevisti quando serve.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Per la Stabilità emotiva, un punteggio alto significa che resti abbastanza tranquillo sotto pressione, non ti lasci buttare giù troppo dalle battute d'arresto e mantieni la calma.`,
				"low":    `Per la Stabilità emotiva, un punteggio basso significa che senti le cose profondamente, cogli le vibrazioni intorno a te e hai una forte consapevolezza emotiva.`,
				"medium": `Per la Stabilità emotiva, un punteggio medio significa che puoi sentire le cose intensamente ma sai anche come gestire quei sentimenti.`,
			},
			domain.TraitOpenness: {
				"high":   `Per l'Apertura mentale, un punteggio alto significa che sei curioso di provare cose nuove, ami pensare creativamente e ti entusiasmi per le grandi idee.`,
				"low":    `Per l'Apertura mentale, un punteggio basso significa che sei pratico e con i piedi per terra, apprezzi ciò che è provato e vero e ti piacciono le routine familiari.`,
				"medium": `Per l'Apertura mentale, un punteggio medio significa che ti piacciono le nuove esperienze ma sai anche quando attenerti a ciò che funziona.`,
			},
		}

	case "uk":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Екстраверсії високий показник означає, що ти заряджаєшся від спілкування з людьми, тобі подобається бути в центрі уваги, і ти почуваєшся як вдома в компаніях.`,
				"low":    `При Екстраверсії низький показник означає, що ти відновлюєшся в тихі моменти наодинці і надаєш перевагу справжнім розмовам з кількома близькими людьми, а не великим тусовкам.`,
				"medium": `При Екстраверсії середній показник означає, що ти можеш перемикатися між спілкуванням з людьми і своїм особистим часом — як підходить моменту.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Доброзичливості високий показник означає, що тобі важливо зберігати мир, ти любиш працювати з іншими і часто думаєш про те, що потрібно оточуючим.`,
				"low":    `При Доброзичливості низький показник означає, що ти кажеш що думаєш, не відступаєш у суперечках і не дозволяєш іншим легко на тебе впливати.`,
				"medium": `При Доброзичливості середній показник означає, що ти можеш бути командним гравцем, коли це важливо, і відстоювати себе, коли треба.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Сумлінності високий показник означає, що тобі подобається бути організованим, робити справи вчасно і триматися довгострокових цілей.`,
				"low":    `При Сумлінності низький показник означає, що ти гнучкий і пливеш за течією, не паришся сильно через плани і легко адаптуєшся, коли щось змінюється.`,
				"medium": `При Сумлінності середній показник означає, що ти можеш бути організованим, коли це важливо, але й справлятися з несподіванками, коли треба.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Емоційній стабільності високий показник означає, що ти залишаєшся досить спокійним під тиском, не сильно вибиваєшся з колії невдачами і зберігаєш холоднокровність.`,
				"low":    `При Емоційній стабільності низький показник означає, що ти глибоко відчуваєш, вловлюєш настрій навколо і маєш сильне емоційне сприйняття.`,
				"medium": `При Емоційній стабільності середній показник означає, що ти можеш відчувати інтенсивно, але також знаєш, як керувати цими почуттями.`,
			},
			domain.TraitOpenness: {
				"high":   `При Відкритості досвіду високий показник означає, що тобі цікаво пробувати нове, ти любиш мислити креативно і загораєшся від крутих ідей.`,
				"low":    `При Відкритості досвіду низький показник означає, що ти практичний і приземлений, цінуєш перевірене часом і любиш звичні ритми.`,
				"medium": `При Відкритості досвіду середній показник означає, що тобі подобається новий досвід, але ти також знаєш, коли триматися того, що працює.`,
			},
		}

	case "bg":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Екстраверсия висок резултат означава, че се зареждаш от това да си с хора, харесва ти да си в центъра на вниманието и се чувстваш като у дома си в групи.`,
				"low":    `При Екстраверсия нисък резултат означава, че се презареждаш в тихи моменти насаме и предпочиташ истински разговори с няколко близки хора пред големи социални сцени.`,
				"medium": `При Екстраверсия среден резултат означава, че можеш да превключваш между това да си с хора и да се наслаждаваш на собствената си компания – каквото пасва на момента.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Сговорчивост висок резултат означава, че ти пука да запазваш мира, обичаш да работиш с други и често мислиш какво имат нужда другите.`,
				"low":    `При Сговорчивост нисък резултат означава, че казваш каквото мислиш, не отстъпваш от несъгласия и не се оставяш лесно да те влияят други.`,
				"medium": `При Сговорчивост среден резултат означава, че можеш да бъдеш екипен играч, когато е важно, и да се защитаваш, когато трябва.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Съвестност висок резултат означава, че обичаш да си организиран, да правиш нещата навреме и да се придържаш към дългосрочните цели.`,
				"low":    `При Съвестност нисък резултат означава, че си гъвкав и следваш потока, не се стресираш много за планове и се адаптираш лесно, когато нещата се променят.`,
				"medium": `При Съвестност среден резултат означава, че можеш да бъдеш организиран, когато е важно, но и да се справяш с неочакваното, когато трябва.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Емоционална стабилност висок резултат означава, че оставаш доста спокоен под напрежение, не се разклащаш много от неуспехи и запазваш хладнокръвие.`,
				"low":    `При Емоционална стабилност нисък резултат означава, че чувстваш нещата дълбоко, улавяш вайба около теб и имаш силно емоционално осъзнаване.`,
				"medium": `При Емоционална стабилност среден резултат означава, че можеш да чувстваш интензивно, но и знаеш как да управляваш тези чувства.`,
			},
			domain.TraitOpenness: {
				"high":   `При Отвореност висок резултат означава, че си любопитен да опитваш нови неща, обичаш да мислиш креативно и се въодушевяваш от големи идеи.`,
				"low":    `При Отвореност нисък резултат означава, че си практичен и реалистичен, цениш изпитаното и вярното и обичаш познатите рутини.`,
				"medium": `При Отвореност среден резултат означава, че се наслаждаваш на нови преживявания, но и знаеш кога да се придържаш към това, което работи.`,
			},
		}

	default: // German (default)
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Bei Extraversion bedeutet ein hoher Wert, dass du Energie daraus ziehst, mit anderen zusammen zu sein, du gerne im Rampenlicht stehst und dich in Gruppen zuhause fühlst.`,
				"low":    `Bei Extraversion bedeutet ein niedriger Wert, dass du dich in ruhigen Momenten allein auflädst und echte Gespräche mit ein paar engen Leuten großen sozialen Events vorziehst.`,
				"medium": `Bei Extraversion bedeutet ein mittlerer Wert, dass du zwischen Zeit mit anderen und Zeit für dich wechseln kannst – je nachdem, was gerade passt.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Bei Verträglichkeit bedeutet ein hoher Wert, dass dir Harmonie wichtig ist, du gerne mit anderen zusammenarbeitest und oft daran denkst, was andere brauchen.`,
				"low":    `Bei Verträglichkeit bedeutet ein niedriger Wert, dass du sagst, was du denkst, bei Meinungsverschiedenheiten nicht zurückweichst und dich nicht so leicht von anderen beeinflussen lässt.`,
				"medium": `Bei Verträglichkeit bedeutet ein mittlerer Wert, dass du ein Teamplayer sein kannst, wenn es drauf ankommt, und für dich einstehen kannst, wenn es nötig ist.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Bei Gewissenhaftigkeit bedeutet ein hoher Wert, dass du gerne organisiert bist, Dinge rechtzeitig erledigst und bei langfristigen Zielen dranbleibst.`,
				"low":    `Bei Gewissenhaftigkeit bedeutet ein niedriger Wert, dass du flexibel bist und mit dem Flow gehst, dich nicht so sehr wegen Plänen stresst und dich leicht anpasst, wenn sich Dinge ändern.`,
				"medium": `Bei Gewissenhaftigkeit bedeutet ein mittlerer Wert, dass du organisiert sein kannst, wenn es drauf ankommt, aber auch mit Überraschungen umgehen kannst, wenn nötig.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Bei Emotionaler Stabilität bedeutet ein hoher Wert, dass du unter Druck ziemlich entspannt bleibst, dich von Rückschlägen nicht so leicht aus der Bahn werfen lässt und cool bleibst.`,
				"low":    `Bei Emotionaler Stabilität bedeutet ein niedriger Wert, dass du Dinge tief empfindest, Stimmungen um dich herum mitbekommst und ein starkes emotionales Bewusstsein hast.`,
				"medium": `Bei Emotionaler Stabilität bedeutet ein mittlerer Wert, dass du Dinge intensiv fühlen kannst, aber auch weißt, wie du mit diesen Gefühlen umgehst.`,
			},
			domain.TraitOpenness: {
				"high":   `Bei Offenheit bedeutet ein hoher Wert, dass du neugierig bist, Neues auszuprobieren, kreatives Denken liebst und dich für große Ideen begeisterst.`,
				"low":    `Bei Offenheit bedeutet ein niedriger Wert, dass du praktisch und bodenständig bist, Bewährtes schätzt und vertraute Routinen magst.`,
				"medium": `Bei Offenheit bedeutet ein mittlerer Wert, dass du neue Erfahrungen genießt, aber auch weißt, wann du bei dem bleibst, was funktioniert.`,
			},
		}
	}
}
