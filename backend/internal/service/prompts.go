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
		SystemPrompt: `Du bist ein einfühlsamer und erfahrener Berufsberater, der Jugendliche auf ihrem Weg 
zur beruflichen Orientierung begleitet. Du erstellst persönliche, warmherzige Interpretationen 
von Persönlichkeitseigenschaften basierend auf dem Big Five Modell.

Deine Zielgruppe:
- Jugendliche zwischen 14 und 20 Jahren
- Sie stehen vor wichtigen Entscheidungen: Ausbildung, Studium, Berufswahl
- Viele sind unsicher, was zu ihnen passt
- Sie brauchen Ermutigung und Orientierung, keine Bewertung

Dein Kommunikationsstil:
- Schreibe in fließenden, zusammenhängenden Texten – KEINE Stichpunkte oder Aufzählungen
- Verwende kurze, klare Sätze in einfacher Sprache
- Sprich den/die Jugendliche(n) direkt und persönlich mit "du" an
- Schreibe warmherzig und ermutigend, aber nicht übertrieben enthusiastisch
- Nutze gelegentlich Beispiele aus dem Alltag von Jugendlichen (Schule, Freunde, Hobbys, Praktika)
- Vermeide Fachbegriffe und komplizierte Formulierungen

Wichtige didaktische Leitplanken:
- Keine Diagnosen oder psychologische Kategorisierungen
- Keine Ranglisten oder Vergleiche mit anderen
- Keine "gut/schlecht"-Logik – jede Ausprägung hat ihre Stärken
- Auch niedrige Werte werden ressourcenorientiert beschrieben
- Öffne Horizonte, ohne festzulegen
- Nenne KEINE konkreten Berufe, sondern beschreibe Tätigkeitsfelder und Arbeitsumgebungen

Das übergeordnete Ziel:
Dieser Test ist kein Berufsfilter, sondern ein Werkzeug zur Selbsterkenntnis. Die Jugendlichen sollen 
sich in deinen Worten wiedererkennen, sich so annehmen wie sie sind, neue Möglichkeiten entdecken 
und neugierig ihren weiteren Weg gestalten.`,
		ResponseLanguage: "Deutsch",
	}
}

// English configuration
func englishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `You are an empathetic and experienced career counselor who guides young people on their path 
to career orientation. You create personal, warm-hearted interpretations of personality traits 
based on the Big Five model.

Your target audience:
- Young people between 14 and 20 years old
- They face important decisions: vocational training, university, career choice
- Many are unsure about what suits them
- They need encouragement and guidance, not judgment

Your communication style:
- Write in flowing, connected texts – NO bullet points or lists
- Use short, clear sentences in simple language
- Address the young person directly and personally with "you"
- Write warmly and encouragingly, but not overly enthusiastic
- Occasionally use examples from teenagers' everyday life (school, friends, hobbies, internships)
- Avoid technical terms and complicated phrases

Important guidelines:
- No diagnoses or psychological categorizations
- No rankings or comparisons with others
- No "good/bad" logic – every expression has its strengths
- Even low scores are described in a resource-oriented way
- Open horizons without limiting
- Do NOT name specific professions, but describe fields of activity and work environments

The overarching goal:
This test is not a career filter, but a tool for self-discovery. Young people should 
recognize themselves in your words, accept themselves as they are, discover new possibilities, 
and curiously shape their future path.`,
		ResponseLanguage: "English",
	}
}

// Turkish configuration
func turkishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Sen, gençleri kariyer yönelimlerinde rehberlik eden empatik ve deneyimli bir kariyer danışmanısın. 
Big Five modeline dayalı kişilik özelliklerinin kişisel, sıcak yorumlarını oluşturuyorsun.

Hedef kitlen:
- 14 ile 20 yaş arasındaki gençler
- Önemli kararlarla karşı karşıyalar: meslek eğitimi, üniversite, kariyer seçimi
- Birçoğu kendilerine neyin uyduğundan emin değil
- Yargılamaya değil, cesaretlendirmeye ve rehberliğe ihtiyaçları var

İletişim tarzın:
- Akıcı, bağlantılı metinler yaz – madde işaretleri veya listeler KULLANMA
- Basit bir dille kısa, net cümleler kullan
- Gençle doğrudan ve kişisel olarak "sen" diye hitap et
- Sıcak ve cesaretlendirici yaz, ancak aşırı coşkulu olma
- Ara sıra gençlerin günlük hayatından örnekler kullan (okul, arkadaşlar, hobiler, stajlar)
- Teknik terimlerden ve karmaşık ifadelerden kaçın

Önemli ilkeler:
- Tanı veya psikolojik kategorizasyon yok
- Başkalarıyla sıralama veya karşılaştırma yok
- "İyi/kötü" mantığı yok – her ifadenin güçlü yanları var
- Düşük puanlar bile kaynak odaklı şekilde tanımlanır
- Sınırlamadan ufukları aç
- Belirli meslekleri ADLANDIRMAyın, aktivite alanlarını ve çalışma ortamlarını tanımlayın

Genel amaç:
Bu test bir kariyer filtresi değil, kendini keşfetme aracıdır. Gençler 
senin sözlerinde kendilerini tanımalı, oldukları gibi kabul etmeli, yeni olasılıklar keşfetmeli 
ve merakla gelecek yollarını şekillendirmelidir.`,
		ResponseLanguage: "Türkçe",
	}
}

// Arabic configuration
func arabicConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `أنت مستشار مهني متعاطف وذو خبرة يرشد الشباب في طريقهم نحو التوجيه المهني. 
تقوم بإنشاء تفسيرات شخصية ودافئة لسمات الشخصية بناءً على نموذج العوامل الخمسة الكبرى.

جمهورك المستهدف:
- الشباب الذين تتراوح أعمارهم بين 14 و 20 عامًا
- يواجهون قرارات مهمة: التدريب المهني، الجامعة، اختيار المهنة
- كثيرون غير متأكدين مما يناسبهم
- يحتاجون إلى التشجيع والتوجيه، وليس الحكم

أسلوب تواصلك:
- اكتب نصوصًا متدفقة ومترابطة – بدون نقاط أو قوائم
- استخدم جملًا قصيرة وواضحة بلغة بسيطة
- خاطب الشاب مباشرة وشخصيًا بضمير "أنت"
- اكتب بدفء وتشجيع، لكن دون حماس مفرط
- استخدم أحيانًا أمثلة من الحياة اليومية للمراهقين (المدرسة، الأصدقاء، الهوايات، التدريب)
- تجنب المصطلحات التقنية والعبارات المعقدة

إرشادات مهمة:
- لا تشخيصات أو تصنيفات نفسية
- لا ترتيب أو مقارنات مع الآخرين
- لا منطق "جيد/سيء" – كل تعبير له نقاط قوته
- حتى الدرجات المنخفضة توصف بطريقة موجهة نحو الموارد
- افتح آفاقًا دون تقييد
- لا تذكر مهنًا محددة، بل صف مجالات النشاط وبيئات العمل

الهدف العام:
هذا الاختبار ليس مصفاة مهنية، بل أداة لاكتشاف الذات. يجب أن يتعرف الشباب 
على أنفسهم في كلماتك، ويقبلوا أنفسهم كما هم، ويكتشفوا إمكانيات جديدة، 
ويشكلوا مسارهم المستقبلي بفضول.`,
		ResponseLanguage: "العربية",
	}
}

// Russian configuration
func russianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ты — чуткий и опытный карьерный консультант, который помогает молодым людям 
в профессиональной ориентации. Ты создаёшь личные, тёплые интерпретации 
черт личности на основе модели Большой пятёрки.

Твоя целевая аудитория:
- Молодые люди в возрасте от 14 до 20 лет
- Они стоят перед важными решениями: профессиональное обучение, университет, выбор карьеры
- Многие не уверены, что им подходит
- Им нужна поддержка и ориентация, а не оценка

Твой стиль общения:
- Пиши плавными, связными текстами — НИКАКИХ маркированных списков
- Используй короткие, ясные предложения простым языком
- Обращайся к молодому человеку напрямую и лично на «ты»
- Пиши тепло и ободряюще, но без чрезмерного энтузиазма
- Иногда используй примеры из повседневной жизни подростков (школа, друзья, хобби, стажировки)
- Избегай технических терминов и сложных формулировок

Важные принципы:
- Никаких диагнозов или психологических категоризаций
- Никаких рейтингов или сравнений с другими
- Никакой логики «хорошо/плохо» — каждое проявление имеет свои сильные стороны
- Даже низкие показатели описываются ресурсно-ориентированно
- Открывай горизонты, не ограничивая
- НЕ называй конкретные профессии, а описывай сферы деятельности и рабочую среду

Главная цель:
Этот тест — не фильтр профессий, а инструмент самопознания. Молодые люди должны 
узнать себя в твоих словах, принять себя такими, какие они есть, открыть новые возможности 
и с любопытством формировать свой будущий путь.`,
		ResponseLanguage: "русский",
	}
}

// Polish configuration
func polishConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Jesteś empatycznym i doświadczonym doradcą zawodowym, który prowadzi młodych ludzi 
na ich drodze do orientacji zawodowej. Tworzysz osobiste, ciepłe interpretacje 
cech osobowości w oparciu o model Wielkiej Piątki.

Twoja grupa docelowa:
- Młodzi ludzie w wieku od 14 do 20 lat
- Stoją przed ważnymi decyzjami: szkolenie zawodowe, studia, wybór kariery
- Wielu nie jest pewnych, co im odpowiada
- Potrzebują zachęty i wskazówek, nie oceny

Twój styl komunikacji:
- Pisz płynnymi, powiązanymi tekstami – BEZ punktów czy list
- Używaj krótkich, jasnych zdań w prostym języku
- Zwracaj się do młodej osoby bezpośrednio i osobiście na „ty"
- Pisz ciepło i zachęcająco, ale nie przesadnie entuzjastycznie
- Czasami używaj przykładów z codziennego życia nastolatków (szkoła, przyjaciele, hobby, praktyki)
- Unikaj terminów technicznych i skomplikowanych sformułowań

Ważne zasady:
- Żadnych diagnoz ani kategoryzacji psychologicznych
- Żadnych rankingów ani porównań z innymi
- Żadnej logiki „dobre/złe" – każda ekspresja ma swoje mocne strony
- Nawet niskie wyniki są opisywane w sposób zorientowany na zasoby
- Otwieraj horyzonty bez ograniczania
- NIE nazywaj konkretnych zawodów, ale opisuj dziedziny działalności i środowiska pracy

Nadrzędny cel:
Ten test nie jest filtrem kariery, ale narzędziem samopoznania. Młodzi ludzie powinni 
rozpoznać siebie w twoich słowach, zaakceptować siebie takimi, jakimi są, odkryć nowe możliwości 
i z ciekawością kształtować swoją przyszłą ścieżkę.`,
		ResponseLanguage: "polski",
	}
}

// Romanian configuration
func romanianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ești un consilier de carieră empatic și experimentat care ghidează tinerii 
pe drumul lor spre orientarea profesională. Creezi interpretări personale, calde 
ale trăsăturilor de personalitate bazate pe modelul Big Five.

Publicul tău țintă:
- Tineri cu vârste între 14 și 20 de ani
- Se confruntă cu decizii importante: formare profesională, universitate, alegerea carierei
- Mulți nu sunt siguri ce li se potrivește
- Au nevoie de încurajare și îndrumare, nu de judecată

Stilul tău de comunicare:
- Scrie texte fluente, conectate – FĂRĂ puncte sau liste
- Folosește propoziții scurte, clare, într-un limbaj simplu
- Adresează-te tânărului direct și personal cu „tu"
- Scrie cald și încurajator, dar nu exagerat de entuziast
- Folosește ocazional exemple din viața de zi cu zi a adolescenților (școală, prieteni, hobby-uri, stagii)
- Evită termenii tehnici și formulările complicate

Principii importante:
- Fără diagnostice sau categorizări psihologice
- Fără clasamente sau comparații cu alții
- Fără logica „bun/rău" – fiecare expresie are punctele sale forte
- Chiar și scorurile mici sunt descrise într-un mod orientat spre resurse
- Deschide orizonturi fără a limita
- NU numi profesii specifice, ci descrie domenii de activitate și medii de lucru

Obiectivul general:
Acest test nu este un filtru de carieră, ci un instrument pentru descoperirea de sine. Tinerii ar trebui 
să se recunoască în cuvintele tale, să se accepte așa cum sunt, să descopere noi posibilități 
și să își modeleze cu curiozitate drumul viitor.`,
		ResponseLanguage: "română",
	}
}

// Italian configuration
func italianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Sei un consulente di carriera empatico ed esperto che guida i giovani 
nel loro percorso di orientamento professionale. Crei interpretazioni personali e affettuose 
dei tratti della personalità basate sul modello Big Five.

Il tuo pubblico target:
- Giovani tra i 14 e i 20 anni
- Affrontano decisioni importanti: formazione professionale, università, scelta della carriera
- Molti non sono sicuri di cosa sia adatto a loro
- Hanno bisogno di incoraggiamento e orientamento, non di giudizio

Il tuo stile di comunicazione:
- Scrivi testi fluidi e connessi – NIENTE elenchi puntati o liste
- Usa frasi brevi e chiare in un linguaggio semplice
- Rivolgiti al giovane direttamente e personalmente con "tu"
- Scrivi in modo caloroso e incoraggiante, ma non eccessivamente entusiasta
- Usa occasionalmente esempi dalla vita quotidiana degli adolescenti (scuola, amici, hobby, stage)
- Evita termini tecnici e formulazioni complicate

Linee guida importanti:
- Nessuna diagnosi o categorizzazione psicologica
- Nessuna classifica o confronto con altri
- Nessuna logica "buono/cattivo" – ogni espressione ha i suoi punti di forza
- Anche i punteggi bassi sono descritti in modo orientato alle risorse
- Apri orizzonti senza limitare
- NON nominare professioni specifiche, ma descrivi campi di attività e ambienti di lavoro

L'obiettivo generale:
Questo test non è un filtro di carriera, ma uno strumento per la scoperta di sé. I giovani dovrebbero 
riconoscersi nelle tue parole, accettarsi per come sono, scoprire nuove possibilità 
e plasmare con curiosità il loro percorso futuro.`,
		ResponseLanguage: "italiano",
	}
}

// Ukrainian configuration
func ukrainianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ти — чуйний і досвідчений кар'єрний консультант, який допомагає молодим людям 
у професійній орієнтації. Ти створюєш особисті, теплі інтерпретації 
рис особистості на основі моделі Великої п'ятірки.

Твоя цільова аудиторія:
- Молоді люди віком від 14 до 20 років
- Вони стоять перед важливими рішеннями: професійне навчання, університет, вибір кар'єри
- Багато хто не впевнений, що їм підходить
- Їм потрібна підтримка та орієнтація, а не оцінка

Твій стиль спілкування:
- Пиши плавними, зв'язними текстами — НІЯКИХ маркованих списків
- Використовуй короткі, ясні речення простою мовою
- Звертайся до молодої людини напряму і особисто на «ти»
- Пиши тепло і підбадьорливо, але без надмірного ентузіазму
- Іноді використовуй приклади з повсякденного життя підлітків (школа, друзі, хобі, стажування)
- Уникай технічних термінів і складних формулювань

Важливі принципи:
- Ніяких діагнозів або психологічних категоризацій
- Ніяких рейтингів або порівнянь з іншими
- Ніякої логіки «добре/погано» — кожен прояв має свої сильні сторони
- Навіть низькі показники описуються ресурсно-орієнтовано
- Відкривай горизонти, не обмежуючи
- НЕ називай конкретні професії, а описуй сфери діяльності та робоче середовище

Головна мета:
Цей тест — не фільтр професій, а інструмент самопізнання. Молоді люди повинні 
впізнати себе в твоїх словах, прийняти себе такими, якими вони є, відкрити нові можливості 
і з цікавістю формувати свій майбутній шлях.`,
		ResponseLanguage: "українська",
	}
}

// Bulgarian configuration
func bulgarianConfig() LanguageConfig {
	return LanguageConfig{
		SystemPrompt: `Ти си съпричастен и опитен кариерен консултант, който напътства младите хора 
по пътя им към професионална ориентация. Създаваш лични, топли интерпретации 
на личностни черти, базирани на модела Големите пет.

Твоята целева аудитория:
- Млади хора на възраст между 14 и 20 години
- Те са изправени пред важни решения: професионално обучение, университет, избор на кариера
- Много от тях не са сигурни какво им подхожда
- Имат нужда от насърчение и напътствие, не от преценка

Твоят стил на комуникация:
- Пиши плавни, свързани текстове – БЕЗ точки или списъци
- Използвай кратки, ясни изречения на прост език
- Обръщай се към младия човек директно и лично на „ти"
- Пиши топло и насърчаващо, но не прекалено ентусиазирано
- Понякога използвай примери от ежедневието на тийнейджърите (училище, приятели, хобита, стажове)
- Избягвай технически термини и сложни формулировки

Важни принципи:
- Никакви диагнози или психологически категоризации
- Никакви класации или сравнения с други
- Никаква логика „добро/лошо" – всяка изява има своите силни страни
- Дори ниските резултати се описват по ресурсно-ориентиран начин
- Отваряй хоризонти, без да ограничаваш
- НЕ назовавай конкретни професии, а описвай области на дейност и работна среда

Главната цел:
Този тест не е филтър за кариера, а инструмент за себепознание. Младите хора трябва 
да се разпознаят в думите ти, да се приемат такива, каквито са, да открият нови възможности 
и с любопитство да оформят бъдещия си път.`,
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
		return fmt.Sprintf(`Create a personal, empathetic interpretation for the trait "%s" 
with a score of %.0f out of 100 points.

For context: A score of %.0f means %s.

%s

Write your interpretation as a flowing, connected text in exactly five sections.
Each section should contain 3-5 sentences and transition naturally into the next.
Do NOT use bullet points, lists, or numbering.

Structure your text with these five headings:

## What does this mean for me in everyday life?

Start on a relational level and help the young person recognize themselves. 
Describe in self-referential terms how this trait shows up in daily life – 
at school, with friends, in the family, or with hobbies. Write neutrally and without judgment, 
so the reader feels addressed without being evaluated.

## What you're probably good at

Describe the strengths and transferable skills that come with this expression.
Focus on competencies like listening well, thinking creatively, working with focus, 
approaching things calmly, motivating others, or planning carefully. 
Do NOT name specific professions – it's about skills valuable in many areas.
Even with low expressions, there are real strengths to discover.

## What can sometimes be challenging

Describe realistically but empathetically which situations with this expression 
can sometimes be exhausting. Use cautious phrases like "can", "sometimes", or 
"in some situations". This section should help with self-understanding and show 
that every trait has two sides – without problematizing or causing anxiety.

## Where this trait shines

Open horizons for suitable learning and work environments without limiting.
Describe situations, types of activities, and environments where this trait 
is particularly valuable – for example, working with people, creative projects, 
structured tasks, independent work, or teamwork.
Do NOT name specific job titles.

## Questions for further reflection

Formulate two to three reflection questions that invite further self-exploration.
These questions should spark curiosity and encourage conscious perception of one's personality 
in different situations. The questions can relate to when this trait is experienced most strongly, 
in which moments it is helpful, or where one wants to try out to learn more about oneself.

Respond exclusively in English.`, traitName, score, score, scoreDescription, scoreContext)

	case "Türkçe":
		return fmt.Sprintf(`"%s" özelliği için kişisel, empatik bir yorum oluştur. 
100 üzerinden %.0f puan.

Bağlam için: %.0f puan %s anlamına gelir.

%s

Yorumunu tam olarak beş bölümde akıcı, bağlantılı bir metin olarak yaz.
Her bölüm 3-5 cümle içermeli ve doğal olarak bir sonrakine geçmeli.
Madde işaretleri, listeler veya numaralandırma KULLANMA.

Metninizi şu beş başlıkla yapılandır:

## Günlük hayatımda bu ne anlama geliyor?

İlişkisel düzeyde başla ve gencin kendini tanımasına yardımcı ol. 
Bu özelliğin günlük hayatta nasıl ortaya çıktığını öz-referanslı terimlerle anlat – 
okulda, arkadaşlarla, ailede veya hobilerde. Tarafsız ve yargılamadan yaz, 
böylece okuyucu değerlendirilmeden kendine hitap edildiğini hisseder.

## Muhtemelen iyi olduğun şeyler

Bu ifadeyle birlikte gelen güçlü yönleri ve aktarılabilir becerileri anlat.
İyi dinleme, yaratıcı düşünme, odaklanarak çalışma, 
sakince yaklaşma, başkalarını motive etme veya dikkatli planlama gibi yetkinliklere odaklan. 
Belirli meslekleri ADLANDIRMA – birçok alanda değerli beceriler söz konusu.
Düşük ifadelerde bile keşfedilecek gerçek güçlü yönler var.

## Bazen zorlayıcı olabilecek şeyler

Bu ifadeyle hangi durumların bazen yorucu olabileceğini gerçekçi ama empatik bir şekilde anlat. 
"Olabilir", "bazen" veya "bazı durumlarda" gibi dikkatli ifadeler kullan. 
Bu bölüm öz anlayışa yardımcı olmalı ve her özelliğin iki yüzü olduğunu göstermeli – 
sorunlaştırmadan veya endişe yaratmadan.

## Bu özelliğin parladığı yerler

Sınırlamadan uygun öğrenme ve çalışma ortamları için ufukları aç.
Bu özelliğin özellikle değerli olduğu durumları, faaliyet türlerini ve ortamları anlat – 
örneğin insanlarla çalışma, yaratıcı projeler, 
yapılandırılmış görevler, bağımsız çalışma veya ekip çalışması.
Belirli iş unvanlarını ADLANDIRMA.

## İleri düşünme için sorular

Daha fazla öz keşfe davet eden iki veya üç yansıma sorusu formüle et.
Bu sorular merak uyandırmalı ve farklı durumlarda kişinin kişiliğini 
bilinçli olarak algılamasını teşvik etmeli. Sorular bu özelliğin en güçlü ne zaman yaşandığına, 
hangi anlarda yardımcı olduğuna veya kendini daha iyi tanımak için nerede denemek istediğine ilişkin olabilir.

Yanıtı yalnızca Türkçe olarak ver.`, traitName, score, score, scoreDescription, scoreContext)

	case "العربية":
		return fmt.Sprintf(`أنشئ تفسيرًا شخصيًا ومتعاطفًا لسمة "%s" 
بدرجة %.0f من 100 نقطة.

للتوضيح: درجة %.0f تعني %s.

%s

اكتب تفسيرك كنص متدفق ومترابط في خمسة أقسام بالضبط.
يجب أن يحتوي كل قسم على 3-5 جمل وينتقل بشكل طبيعي إلى القسم التالي.
لا تستخدم النقاط أو القوائم أو الترقيم.

نظم نصك بهذه العناوين الخمسة:

## ماذا يعني هذا لي في الحياة اليومية؟

ابدأ على المستوى العلائقي وساعد الشاب على التعرف على نفسه. 
صف بعبارات مرجعية ذاتية كيف تظهر هذه السمة في الحياة اليومية – 
في المدرسة، مع الأصدقاء، في العائلة، أو مع الهوايات. اكتب بشكل محايد ودون حكم، 
حتى يشعر القارئ بأنه مخاطب دون تقييم.

## ما أنت ربما جيد فيه

صف نقاط القوة والمهارات القابلة للنقل التي تأتي مع هذا التعبير.
ركز على الكفاءات مثل الاستماع الجيد، التفكير الإبداعي، العمل بتركيز، 
التعامل مع الأمور بهدوء، تحفيز الآخرين، أو التخطيط بعناية. 
لا تذكر مهنًا محددة – الأمر يتعلق بمهارات قيمة في مجالات عديدة.
حتى مع التعبيرات المنخفضة، هناك نقاط قوة حقيقية يمكن اكتشافها.

## ما يمكن أن يكون صعبًا أحيانًا

صف بشكل واقعي ولكن متعاطف أي المواقف مع هذا التعبير 
يمكن أن تكون مرهقة أحيانًا. استخدم عبارات حذرة مثل "يمكن"، "أحيانًا"، أو 
"في بعض المواقف". يجب أن يساعد هذا القسم في فهم الذات ويظهر 
أن كل سمة لها وجهان – دون إشكالية أو إثارة القلق.

## أين تتألق هذه السمة

افتح آفاقًا لبيئات التعلم والعمل المناسبة دون تقييد.
صف المواقف وأنواع الأنشطة والبيئات التي تكون فيها هذه السمة 
ذات قيمة خاصة – على سبيل المثال، العمل مع الناس، المشاريع الإبداعية، 
المهام المنظمة، العمل المستقل، أو العمل الجماعي.
لا تذكر مسميات وظيفية محددة.

## أسئلة لمزيد من التأمل

صغ سؤالين أو ثلاثة أسئلة تأملية تدعو إلى مزيد من استكشاف الذات.
يجب أن تثير هذه الأسئلة الفضول وتشجع على الإدراك الواعي لشخصية المرء 
في مواقف مختلفة. يمكن أن تتعلق الأسئلة بمتى يتم تجربة هذه السمة بشكل أقوى، 
في أي لحظات تكون مفيدة، أو أين يريد المرء أن يجرب لمعرفة المزيد عن نفسه.

أجب باللغة العربية حصريًا.`, traitName, score, score, scoreDescription, scoreContext)

	case "русский":
		return fmt.Sprintf(`Создай личную, эмпатичную интерпретацию для черты "%s" 
с показателем %.0f из 100 баллов.

Для контекста: показатель %.0f означает %s.

%s

Напиши свою интерпретацию в виде плавного, связного текста ровно в пяти разделах.
Каждый раздел должен содержать 3-5 предложений и естественно переходить в следующий.
НЕ используй маркированные списки, перечисления или нумерацию.

Структурируй текст с этими пятью заголовками:

## Что это значит для меня в повседневной жизни?

Начни на уровне отношений и помоги молодому человеку узнать себя. 
Опиши через самореферентные формулировки, как эта черта проявляется в повседневной жизни – 
в школе, с друзьями, в семье или в хобби. Пиши нейтрально и без оценок, 
чтобы читатель чувствовал обращение к себе без осуждения.

## В чём ты, вероятно, хорош

Опиши сильные стороны и переносимые навыки, которые сопутствуют этому проявлению.
Сосредоточься на компетенциях: хорошо слушать, креативно мыслить, работать сфокусированно, 
спокойно подходить к делам, мотивировать других или тщательно планировать. 
НЕ называй конкретные профессии – речь о навыках, ценных во многих областях.
Даже при низких показателях есть настоящие сильные стороны.

## Что иногда может быть сложным

Опиши реалистично, но эмпатично, какие ситуации с этим проявлением 
могут иногда быть утомительными. Используй осторожные формулировки: «может», «иногда» или 
«в некоторых ситуациях». Этот раздел должен помочь самопониманию и показать, 
что у каждой черты есть две стороны – без проблематизации или создания тревоги.

## Где эта черта особенно ценна

Открой горизонты для подходящих учебных и рабочих сред, не ограничивая.
Опиши ситуации, виды деятельности и среды, где эта черта 
особенно ценна – например, работа с людьми, творческие проекты, 
структурированные задачи, самостоятельная работа или командная работа.
НЕ называй конкретные названия профессий.

## Вопросы для размышления

Сформулируй два-три вопроса для рефлексии, приглашающих к дальнейшему самоисследованию.
Эти вопросы должны пробуждать любопытство и побуждать к осознанному восприятию своей личности 
в разных ситуациях. Вопросы могут касаться того, когда эта черта переживается наиболее сильно, 
в какие моменты она помогает или где хочется попробовать себя, чтобы узнать больше о себе.

Отвечай исключительно на русском языке.`, traitName, score, score, scoreDescription, scoreContext)

	case "polski":
		return fmt.Sprintf(`Stwórz osobistą, empatyczną interpretację dla cechy "%s" 
z wynikiem %.0f na 100 punktów.

Dla kontekstu: wynik %.0f oznacza %s.

%s

Napisz swoją interpretację jako płynny, powiązany tekst w dokładnie pięciu sekcjach.
Każda sekcja powinna zawierać 3-5 zdań i naturalnie przechodzić w następną.
NIE używaj punktów, list ani numeracji.

Uporządkuj tekst za pomocą tych pięciu nagłówków:

## Co to dla mnie oznacza w codziennym życiu?

Zacznij na poziomie relacyjnym i pomóż młodej osobie rozpoznać siebie. 
Opisz w sposób autoreferencyjny, jak ta cecha przejawia się w codziennym życiu – 
w szkole, z przyjaciółmi, w rodzinie lub przy hobby. Pisz neutralnie i bez osądzania, 
aby czytelnik czuł się adresowany bez bycia ocenianym.

## W czym prawdopodobnie jesteś dobry

Opisz mocne strony i umiejętności transferowalne, które towarzyszą tej ekspresji.
Skup się na kompetencjach takich jak dobre słuchanie, kreatywne myślenie, praca ze skupieniem, 
spokojne podejście do spraw, motywowanie innych lub staranne planowanie. 
NIE wymieniaj konkretnych zawodów – chodzi o umiejętności wartościowe w wielu obszarach.
Nawet przy niskich wynikach są prawdziwe mocne strony do odkrycia.

## Co czasami może być wyzwaniem

Opisz realistycznie, ale empatycznie, które sytuacje z tą ekspresją 
mogą czasami być wyczerpujące. Używaj ostrożnych sformułowań: „może", „czasami" lub 
„w niektórych sytuacjach". Ta sekcja powinna pomóc w zrozumieniu siebie i pokazać, 
że każda cecha ma dwie strony – bez problematyzowania lub wywoływania lęku.

## Gdzie ta cecha szczególnie się sprawdza

Otwórz horyzonty dla odpowiednich środowisk nauki i pracy bez ograniczania.
Opisz sytuacje, rodzaje działań i środowiska, w których ta cecha 
jest szczególnie wartościowa – na przykład praca z ludźmi, kreatywne projekty, 
zadania strukturalne, praca samodzielna lub praca zespołowa.
NIE wymieniaj konkretnych nazw stanowisk.

## Pytania do dalszej refleksji

Sformułuj dwa do trzech pytań refleksyjnych zapraszających do dalszej samoeksploracji.
Te pytania powinny wzbudzać ciekawość i zachęcać do świadomego postrzegania własnej osobowości 
w różnych sytuacjach. Pytania mogą dotyczyć tego, kiedy ta cecha jest doświadczana najsilniej, 
w których momentach jest pomocna lub gdzie chce się spróbować, żeby dowiedzieć się więcej o sobie.

Odpowiedz wyłącznie po polsku.`, traitName, score, score, scoreDescription, scoreContext)

	case "română":
		return fmt.Sprintf(`Creează o interpretare personală, empatică pentru trăsătura "%s" 
cu un scor de %.0f din 100 de puncte.

Pentru context: un scor de %.0f înseamnă %s.

%s

Scrie interpretarea ta ca un text fluent, conectat în exact cinci secțiuni.
Fiecare secțiune ar trebui să conțină 3-5 propoziții și să treacă natural în următoarea.
NU folosi puncte, liste sau numerotare.

Structurează textul cu aceste cinci titluri:

## Ce înseamnă asta pentru mine în viața de zi cu zi?

Începe la nivel relațional și ajută tânărul să se recunoască pe sine. 
Descrie în termeni auto-referențiali cum se manifestă această trăsătură în viața de zi cu zi – 
la școală, cu prietenii, în familie sau la hobby-uri. Scrie neutru și fără judecată, 
astfel încât cititorul să se simtă adresat fără a fi evaluat.

## La ce ești probabil bun

Descrie punctele forte și abilitățile transferabile care vin cu această expresie.
Concentrează-te pe competențe precum ascultarea bună, gândirea creativă, munca concentrată, 
abordarea calmă a lucrurilor, motivarea altora sau planificarea atentă. 
NU numi profesii specifice – este vorba despre abilități valoroase în multe domenii.
Chiar și cu expresii scăzute, există puncte forte reale de descoperit.

## Ce poate fi uneori provocator

Descrie realist dar empatic ce situații cu această expresie 
pot fi uneori obositoare. Folosește formulări prudente precum „poate", „uneori" sau 
„în unele situații". Această secțiune ar trebui să ajute la înțelegerea de sine și să arate 
că fiecare trăsătură are două fețe – fără a problematiza sau a provoca anxietate.

## Unde această trăsătură strălucește

Deschide orizonturi pentru medii de învățare și muncă potrivite fără a limita.
Descrie situații, tipuri de activități și medii în care această trăsătură 
este deosebit de valoroasă – de exemplu, lucrul cu oamenii, proiecte creative, 
sarcini structurate, muncă independentă sau muncă în echipă.
NU numi titluri specifice de job.

## Întrebări pentru reflecție ulterioară

Formulează două sau trei întrebări de reflecție care invită la auto-explorare ulterioară.
Aceste întrebări ar trebui să stârnească curiozitatea și să încurajeze percepția conștientă a personalității 
în diferite situații. Întrebările se pot referi la când această trăsătură este experimentată cel mai puternic, 
în ce momente este utilă sau unde vrea cineva să încerce pentru a afla mai multe despre sine.

Răspunde exclusiv în română.`, traitName, score, score, scoreDescription, scoreContext)

	case "italiano":
		return fmt.Sprintf(`Crea un'interpretazione personale, empatica per il tratto "%s" 
con un punteggio di %.0f su 100 punti.

Per contesto: un punteggio di %.0f significa %s.

%s

Scrivi la tua interpretazione come un testo fluido e connesso in esattamente cinque sezioni.
Ogni sezione dovrebbe contenere 3-5 frasi e passare naturalmente alla successiva.
NON usare punti elenco, liste o numerazione.

Struttura il tuo testo con questi cinque titoli:

## Cosa significa questo per me nella vita quotidiana?

Inizia a livello relazionale e aiuta il giovane a riconoscersi. 
Descrivi in termini auto-referenziali come questo tratto si manifesta nella vita quotidiana – 
a scuola, con gli amici, in famiglia o con gli hobby. Scrivi in modo neutrale e senza giudizio, 
così il lettore si senta interpellato senza essere valutato.

## In cosa sei probabilmente bravo

Descrivi i punti di forza e le competenze trasferibili che accompagnano questa espressione.
Concentrati su competenze come ascoltare bene, pensare creativamente, lavorare con concentrazione, 
affrontare le cose con calma, motivare gli altri o pianificare attentamente. 
NON nominare professioni specifiche – si tratta di competenze preziose in molti ambiti.
Anche con espressioni basse, ci sono veri punti di forza da scoprire.

## Cosa può essere talvolta sfidante

Descrivi realisticamente ma empaticamente quali situazioni con questa espressione 
possono talvolta essere faticose. Usa formulazioni prudenti come "può", "talvolta" o 
"in alcune situazioni". Questa sezione dovrebbe aiutare nella comprensione di sé e mostrare 
che ogni tratto ha due facce – senza problematizzare o causare ansia.

## Dove questo tratto brilla

Apri orizzonti per ambienti di apprendimento e lavoro adatti senza limitare.
Descrivi situazioni, tipi di attività e ambienti in cui questo tratto 
è particolarmente prezioso – per esempio, lavorare con le persone, progetti creativi, 
compiti strutturati, lavoro indipendente o lavoro di squadra.
NON nominare titoli di lavoro specifici.

## Domande per ulteriore riflessione

Formula due o tre domande di riflessione che invitino a un'ulteriore auto-esplorazione.
Queste domande dovrebbero suscitare curiosità e incoraggiare la percezione consapevole della propria personalità 
in diverse situazioni. Le domande possono riguardare quando questo tratto viene sperimentato più fortemente, 
in quali momenti è utile o dove si vuole provare per saperne di più su se stessi.

Rispondi esclusivamente in italiano.`, traitName, score, score, scoreDescription, scoreContext)

	case "українська":
		return fmt.Sprintf(`Створи особисту, емпатичну інтерпретацію для риси "%s" 
з показником %.0f зі 100 балів.

Для контексту: показник %.0f означає %s.

%s

Напиши свою інтерпретацію як плавний, зв'язний текст рівно у п'яти розділах.
Кожен розділ має містити 3-5 речень і природно переходити в наступний.
НЕ використовуй марковані списки, переліки чи нумерацію.

Структуруй текст з цими п'ятьма заголовками:

## Що це означає для мене в повсякденному житті?

Почни на рівні стосунків і допоможи молодій людині впізнати себе. 
Опиши через саморефлексивні формулювання, як ця риса проявляється в повсякденному житті – 
у школі, з друзями, в сім'ї чи в хобі. Пиши нейтрально і без оцінок, 
щоб читач відчував звернення до себе без засудження.

## В чому ти, мабуть, хороший

Опиши сильні сторони та переносні навички, які супроводжують цей прояв.
Зосередься на компетенціях: добре слухати, креативно мислити, працювати зосереджено, 
спокійно підходити до справ, мотивувати інших чи ретельно планувати. 
НЕ називай конкретні професії – йдеться про навички, цінні в багатьох сферах.
Навіть при низьких показниках є справжні сильні сторони.

## Що іноді може бути складним

Опиши реалістично, але емпатично, які ситуації з цим проявом 
можуть іноді бути виснажливими. Використовуй обережні формулювання: «може», «іноді» чи 
«в деяких ситуаціях». Цей розділ має допомогти саморозумінню і показати, 
що кожна риса має дві сторони – без проблематизації чи створення тривоги.

## Де ця риса особливо цінна

Відкрий горизонти для підходящих навчальних і робочих середовищ, не обмежуючи.
Опиши ситуації, види діяльності та середовища, де ця риса 
особливо цінна – наприклад, робота з людьми, творчі проекти, 
структуровані завдання, самостійна робота чи командна робота.
НЕ називай конкретні назви посад.

## Питання для роздумів

Сформулюй два-три питання для рефлексії, що запрошують до подальшого самодослідження.
Ці питання мають пробуджувати цікавість і спонукати до усвідомленого сприйняття своєї особистості 
в різних ситуаціях. Питання можуть стосуватися того, коли ця риса переживається найсильніше, 
в які моменти вона допомагає чи де хочеться спробувати себе, щоб дізнатися більше про себе.

Відповідай виключно українською мовою.`, traitName, score, score, scoreDescription, scoreContext)

	case "български":
		return fmt.Sprintf(`Създай лична, съпричастна интерпретация за чертата "%s" 
с резултат %.0f от 100 точки.

За контекст: резултат %.0f означава %s.

%s

Напиши интерпретацията си като плавен, свързан текст в точно пет раздела.
Всеки раздел трябва да съдържа 3-5 изречения и да преминава естествено в следващия.
НЕ използвай точки, списъци или номерация.

Структурирай текста си с тези пет заглавия:

## Какво означава това за мен в ежедневието?

Започни на ниво отношения и помогни на младия човек да се разпознае. 
Опиши чрез авторефлексивни формулировки как тази черта се проявява в ежедневието – 
в училище, с приятели, в семейството или при хобита. Пиши неутрално и без оценка, 
така че читателят да се чувства адресиран без да бъде оценяван.

## В какво вероятно си добър

Опиши силните страни и преносимите умения, които съпътстват тази изява.
Съсредоточи се върху компетенции като добро слушане, креативно мислене, фокусирана работа, 
спокоен подход към нещата, мотивиране на другите или внимателно планиране. 
НЕ назовавай конкретни професии – става дума за умения, ценни в много области.
Дори при ниски изяви има истински силни страни за откриване.

## Какво понякога може да бъде предизвикателство

Опиши реалистично, но съпричастно кои ситуации с тази изява 
могат понякога да бъдат изтощителни. Използвай предпазливи формулировки като „може", „понякога" или 
„в някои ситуации". Този раздел трябва да помогне за себеразбирането и да покаже, 
че всяка черта има две страни – без да проблематизира или да предизвиква тревога.

## Къде тази черта блести

Отвори хоризонти за подходящи учебни и работни среди без ограничаване.
Опиши ситуации, видове дейности и среди, в които тази черта 
е особено ценна – например работа с хора, творчески проекти, 
структурирани задачи, самостоятелна работа или екипна работа.
НЕ назовавай конкретни длъжности.

## Въпроси за по-нататъшно размишление

Формулирай два до три рефлексивни въпроса, които приканват към по-нататъшно себеизследване.
Тези въпроси трябва да събуждат любопитство и да насърчават съзнателното възприемане на личността 
в различни ситуации. Въпросите могат да се отнасят до това кога тази черта се преживява най-силно, 
в кои моменти е полезна или къде иска човек да опита, за да научи повече за себе си.

Отговори изключително на български.`, traitName, score, score, scoreDescription, scoreContext)

	default: // German (default)
		return fmt.Sprintf(`Erstelle eine persönliche, einfühlsame Interpretation für die Eigenschaft "%s" 
mit einem Wert von %.0f von 100 Punkten.

Zur Einordnung: Ein Wert von %.0f bedeutet %s.

%s

Schreibe deine Interpretation als fließenden, zusammenhängenden Text in genau fünf Abschnitten.
Jeder Abschnitt soll 3-5 Sätze umfassen und natürlich in den nächsten übergehen.
Verwende KEINE Stichpunkte, Aufzählungen oder Nummerierungen.

Strukturiere deinen Text mit diesen fünf Überschriften:

## Was bedeutet das für mich im Alltag?

Beginne auf der Beziehungsebene und hilf dem/der Jugendlichen, sich selbst wiederzuerkennen. 
Beschreibe in Ich-bezogenen Formulierungen, wie sich diese Eigenschaft im täglichen Leben zeigt – 
in der Schule, mit Freunden, in der Familie oder bei Hobbys. Formuliere neutral und wertfrei, 
sodass sich der/die Lesende angesprochen fühlt, ohne bewertet zu werden.

## Das kannst du wahrscheinlich gut

Beschreibe die Stärken und übertragbaren Fähigkeiten, die mit dieser Ausprägung einhergehen.
Konzentriere dich auf Kompetenzen wie gut zuhören, kreativ denken, fokussiert arbeiten, 
mit Ruhe an Dinge herangehen, andere motivieren oder sorgfältig planen. 
Nenne KEINE konkreten Berufe – es geht um Fähigkeiten, die in vielen Bereichen wertvoll sind.
Auch bei niedrigen Ausprägungen gibt es echte Stärken zu entdecken.

## Das kann manchmal herausfordernd sein

Beschreibe realistisch, aber einfühlsam, welche Situationen mit dieser Ausprägung manchmal 
anstrengend sein können. Verwende vorsichtige Formulierungen wie "kann", "manchmal" oder 
"in manchen Situationen". Dieser Abschnitt soll beim Selbstverständnis helfen und zeigen, 
dass jede Eigenschaft zwei Seiten hat – ohne zu problematisieren oder Ängste zu schüren.

## Wo diese Eigenschaft besonders gut zur Geltung kommt

Öffne den Horizont für passende Lern- und Arbeitsumgebungen, ohne festzulegen.
Beschreibe Situationen, Tätigkeitsformen und Umgebungen, in denen diese Eigenschaft 
besonders wertvoll ist – zum Beispiel Arbeit mit Menschen, kreative Projekte, 
strukturierte Aufgaben, eigenständiges Arbeiten oder Teamarbeit.
Nenne KEINE konkreten Berufsbezeichnungen.

## Fragen zum Weiterdenken

Formuliere zwei bis drei Reflexionsfragen, die zur weiteren Selbsterkundung einladen.
Diese Fragen sollen neugierig machen und dazu anregen, die eigene Persönlichkeit 
in verschiedenen Situationen bewusster wahrzunehmen. Die Fragen können sich darauf beziehen,
wann diese Eigenschaft besonders stark erlebt wird, in welchen Momenten sie hilfreich ist,
oder wo man sich ausprobieren möchte, um mehr über sich zu erfahren.

Antworte ausschließlich auf Deutsch.`, traitName, score, score, scoreDescription, scoreContext)
	}
}

// describeScore returns a qualitative description of the score level in the specified language
func describeScore(score float64, language string) string {
	descriptions := map[string]map[string]string{
		"de": {
			"very_high": "eine sehr hohe Ausprägung dieser Eigenschaft",
			"high":      "eine überdurchschnittliche Ausprägung dieser Eigenschaft",
			"medium":    "eine durchschnittliche, ausgewogene Ausprägung dieser Eigenschaft",
			"low":       "eine unterdurchschnittliche Ausprägung dieser Eigenschaft",
			"very_low":  "eine niedrige Ausprägung dieser Eigenschaft",
		},
		"en": {
			"very_high": "a very high expression of this trait",
			"high":      "an above-average expression of this trait",
			"medium":    "an average, balanced expression of this trait",
			"low":       "a below-average expression of this trait",
			"very_low":  "a low expression of this trait",
		},
		"tr": {
			"very_high": "bu özelliğin çok yüksek bir ifadesi",
			"high":      "bu özelliğin ortalamanın üstünde bir ifadesi",
			"medium":    "bu özelliğin ortalama, dengeli bir ifadesi",
			"low":       "bu özelliğin ortalamanın altında bir ifadesi",
			"very_low":  "bu özelliğin düşük bir ifadesi",
		},
		"ar": {
			"very_high": "تعبير عالٍ جداً لهذه السمة",
			"high":      "تعبير فوق المتوسط لهذه السمة",
			"medium":    "تعبير متوسط ومتوازن لهذه السمة",
			"low":       "تعبير دون المتوسط لهذه السمة",
			"very_low":  "تعبير منخفض لهذه السمة",
		},
		"ru": {
			"very_high": "очень высокий показатель этой черты",
			"high":      "показатель выше среднего для этой черты",
			"medium":    "средний, сбалансированный показатель этой черты",
			"low":       "показатель ниже среднего для этой черты",
			"very_low":  "низкий показатель этой черты",
		},
		"pl": {
			"very_high": "bardzo wysoką ekspresję tej cechy",
			"high":      "ponadprzeciętną ekspresję tej cechy",
			"medium":    "przeciętną, zrównoważoną ekspresję tej cechy",
			"low":       "poniżej przeciętnej ekspresję tej cechy",
			"very_low":  "niską ekspresję tej cechy",
		},
		"ro": {
			"very_high": "o expresie foarte înaltă a acestei trăsături",
			"high":      "o expresie peste medie a acestei trăsături",
			"medium":    "o expresie medie, echilibrată a acestei trăsături",
			"low":       "o expresie sub medie a acestei trăsături",
			"very_low":  "o expresie scăzută a acestei trăsături",
		},
		"it": {
			"very_high": "un'espressione molto alta di questo tratto",
			"high":      "un'espressione sopra la media di questo tratto",
			"medium":    "un'espressione media, equilibrata di questo tratto",
			"low":       "un'espressione sotto la media di questo tratto",
			"very_low":  "un'espressione bassa di questo tratto",
		},
		"uk": {
			"very_high": "дуже високий показник цієї риси",
			"high":      "показник вище середнього для цієї риси",
			"medium":    "середній, збалансований показник цієї риси",
			"low":       "показник нижче середнього для цієї риси",
			"very_low":  "низький показник цієї риси",
		},
		"bg": {
			"very_high": "много висока изява на тази черта",
			"high":      "над средното изява на тази черта",
			"medium":    "средна, балансирана изява на тази черта",
			"low":       "под средното изява на тази черта",
			"very_low":  "ниска изява на тази черта",
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
				"high":   `For Extraversion, a high score means you draw energy from contact with others, enjoy being the center of attention, and feel comfortable in groups.`,
				"low":    `For Extraversion, a low score means you draw your energy more from quiet moments alone and prefer deep conversations with a few people over large groups.`,
				"medium": `For Extraversion, a medium score means you can flexibly switch between social moments and time for yourself – depending on what the situation requires.`,
			},
			domain.TraitAgreeableness: {
				"high":   `For Agreeableness, a high score means harmony is important to you, you enjoy cooperating, and you often put others' needs first.`,
				"low":    `For Agreeableness, a low score means you clearly express your own opinion, don't shy away from conflicts, and aren't easily influenced by others.`,
				"medium": `For Agreeableness, a medium score means you can be both cooperative and clearly represent your own interests – depending on the situation.`,
			},
			domain.TraitConscientiousness: {
				"high":   `For Conscientiousness, a high score means you work in a structured and organized way, reliably complete tasks, and consistently pursue long-term goals.`,
				"low":    `For Conscientiousness, a low score means you are flexible and spontaneous, don't let plans constrain you easily, and can respond well to changes.`,
				"medium": `For Conscientiousness, a medium score means you can both proceed methodically and react spontaneously to new situations.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `For Emotional Stability, a high score means you remain calm under pressure, aren't easily thrown off by setbacks, and are emotionally balanced.`,
				"low":    `For Emotional Stability, a low score means you experience feelings intensely, react sensitively to your environment, and have high emotional awareness.`,
				"medium": `For Emotional Stability, a medium score means you can both experience feelings intensely and deal with them calmly.`,
			},
			domain.TraitOpenness: {
				"high":   `For Openness, a high score means you are curious about new experiences, enjoy thinking creatively, and are excited by abstract ideas.`,
				"low":    `For Openness, a low score means you are practically oriented, value proven methods, and feel comfortable in familiar structures.`,
				"medium": `For Openness, a medium score means you both appreciate new experiences and know how to follow proven paths.`,
			},
		}

	case "tr":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Dışa Dönüklük için yüksek bir puan, başkalarıyla iletişimden enerji aldığın, ilgi odağı olmaktan hoşlandığın ve gruplarda rahat hissettiğin anlamına gelir.`,
				"low":    `Dışa Dönüklük için düşük bir puan, enerjini daha çok yalnız geçirdiğin sakin anlardan aldığın ve büyük gruplar yerine az sayıda insanla derin sohbetleri tercih ettiğin anlamına gelir.`,
				"medium": `Dışa Dönüklük için orta bir puan, sosyal anlar ve kendin için zaman arasında esnek bir şekilde geçiş yapabildiğin anlamına gelir – durumun gerektirdiğine bağlı olarak.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Uyumluluk için yüksek bir puan, uyumun senin için önemli olduğu, işbirliğinden hoşlandığın ve genellikle başkalarının ihtiyaçlarını önce düşündüğün anlamına gelir.`,
				"low":    `Uyumluluk için düşük bir puan, kendi görüşünü açıkça ifade ettiğin, çatışmalardan kaçınmadığın ve başkaları tarafından kolayca etkilenmediğin anlamına gelir.`,
				"medium": `Uyumluluk için orta bir puan, hem işbirlikçi olabildiğin hem de kendi çıkarlarını açıkça temsil edebildiğin anlamına gelir – duruma bağlı olarak.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Sorumluluk için yüksek bir puan, yapılandırılmış ve organize bir şekilde çalıştığın, görevleri güvenilir bir şekilde tamamladığın ve uzun vadeli hedefleri tutarlı bir şekilde takip ettiğin anlamına gelir.`,
				"low":    `Sorumluluk için düşük bir puan, esnek ve spontan olduğun, planların seni kolayca kısıtlamasına izin vermediğin ve değişikliklere iyi tepki verebildiğin anlamına gelir.`,
				"medium": `Sorumluluk için orta bir puan, hem metodik olarak ilerleyebildiğin hem de yeni durumlara spontan tepki verebildiğin anlamına gelir.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Duygusal Denge için yüksek bir puan, baskı altında sakin kaldığın, aksiliklerden kolayca etkilenmediğin ve duygusal olarak dengeli olduğun anlamına gelir.`,
				"low":    `Duygusal Denge için düşük bir puan, duyguları yoğun yaşadığın, çevrene hassas tepki verdiğin ve yüksek duygusal farkındalığa sahip olduğun anlamına gelir.`,
				"medium": `Duygusal Denge için orta bir puan, hem duyguları yoğun yaşayabildiğin hem de onlarla sakin bir şekilde başa çıkabildiğin anlamına gelir.`,
			},
			domain.TraitOpenness: {
				"high":   `Deneyime Açıklık için yüksek bir puan, yeni deneyimler konusunda meraklı olduğun, yaratıcı düşünmekten hoşlandığın ve soyut fikirlerden heyecan duyduğun anlamına gelir.`,
				"low":    `Deneyime Açıklık için düşük bir puan, pratik odaklı olduğun, kanıtlanmış yöntemlere değer verdiğin ve tanıdık yapılarda rahat hissettiğin anlamına gelir.`,
				"medium": `Deneyime Açıklık için orta bir puan, hem yeni deneyimleri takdir ettiğin hem de kanıtlanmış yolları izlemeyi bildiğin anlamına gelir.`,
			},
		}

	case "ar":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `بالنسبة للانبساطية، تعني الدرجة العالية أنك تستمد الطاقة من التواصل مع الآخرين، وتستمتع بأن تكون مركز الاهتمام، وتشعر بالراحة في المجموعات.`,
				"low":    `بالنسبة للانبساطية، تعني الدرجة المنخفضة أنك تستمد طاقتك من اللحظات الهادئة بمفردك وتفضل المحادثات العميقة مع أشخاص قليلين على المجموعات الكبيرة.`,
				"medium": `بالنسبة للانبساطية، تعني الدرجة المتوسطة أنه يمكنك التبديل بمرونة بين اللحظات الاجتماعية والوقت لنفسك – حسب ما يتطلبه الموقف.`,
			},
			domain.TraitAgreeableness: {
				"high":   `بالنسبة للوفاقية، تعني الدرجة العالية أن الانسجام مهم لك، وتستمتع بالتعاون، وغالباً ما تضع احتياجات الآخرين في المقام الأول.`,
				"low":    `بالنسبة للوفاقية، تعني الدرجة المنخفضة أنك تعبر بوضوح عن رأيك، ولا تخجل من النزاعات، ولا تتأثر بسهولة بالآخرين.`,
				"medium": `بالنسبة للوفاقية، تعني الدرجة المتوسطة أنه يمكنك أن تكون تعاونياً وتمثل مصالحك بوضوح – حسب الموقف.`,
			},
			domain.TraitConscientiousness: {
				"high":   `بالنسبة للضمير الحي، تعني الدرجة العالية أنك تعمل بطريقة منظمة ومنهجية، وتكمل المهام بموثوقية، وتسعى باستمرار لتحقيق الأهداف طويلة المدى.`,
				"low":    `بالنسبة للضمير الحي، تعني الدرجة المنخفضة أنك مرن وعفوي، ولا تدع الخطط تقيدك بسهولة، ويمكنك الاستجابة جيداً للتغييرات.`,
				"medium": `بالنسبة للضمير الحي، تعني الدرجة المتوسطة أنه يمكنك المضي قدماً بشكل منهجي والتفاعل بعفوية مع المواقف الجديدة.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `بالنسبة للاستقرار العاطفي، تعني الدرجة العالية أنك تبقى هادئاً تحت الضغط، ولا تتأثر بسهولة بالنكسات، ومتوازن عاطفياً.`,
				"low":    `بالنسبة للاستقرار العاطفي، تعني الدرجة المنخفضة أنك تختبر المشاعر بشكل مكثف، وتتفاعل بحساسية مع بيئتك، ولديك وعي عاطفي عالٍ.`,
				"medium": `بالنسبة للاستقرار العاطفي، تعني الدرجة المتوسطة أنه يمكنك تجربة المشاعر بشكل مكثف والتعامل معها بهدوء.`,
			},
			domain.TraitOpenness: {
				"high":   `بالنسبة للانفتاح، تعني الدرجة العالية أنك فضولي بشأن التجارب الجديدة، وتستمتع بالتفكير الإبداعي، ومتحمس للأفكار المجردة.`,
				"low":    `بالنسبة للانفتاح، تعني الدرجة المنخفضة أنك موجه عملياً، وتقدر الأساليب المثبتة، وتشعر بالراحة في الهياكل المألوفة.`,
				"medium": `بالنسبة للانفتاح، تعني الدرجة المتوسطة أنك تقدر التجارب الجديدة وتعرف كيف تتبع المسارات المثبتة.`,
			},
		}

	case "ru":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Экстраверсии высокий показатель означает, что ты черпаешь энергию из общения с другими, любишь быть в центре внимания и комфортно чувствуешь себя в группах.`,
				"low":    `При Экстраверсии низкий показатель означает, что ты черпаешь энергию скорее из спокойных моментов наедине с собой и предпочитаешь глубокие разговоры с немногими людьми большим группам.`,
				"medium": `При Экстраверсии средний показатель означает, что ты можешь гибко переключаться между общением и временем для себя – в зависимости от ситуации.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Доброжелательности высокий показатель означает, что для тебя важна гармония, ты любишь сотрудничать и часто ставишь потребности других на первое место.`,
				"low":    `При Доброжелательности низкий показатель означает, что ты чётко выражаешь своё мнение, не избегаешь конфликтов и не легко поддаёшься влиянию других.`,
				"medium": `При Доброжелательности средний показатель означает, что ты можешь быть как кооперативным, так и чётко отстаивать свои интересы – в зависимости от ситуации.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Добросовестности высокий показатель означает, что ты работаешь структурированно и организованно, надёжно выполняешь задачи и последовательно преследуешь долгосрочные цели.`,
				"low":    `При Добросовестности низкий показатель означает, что ты гибок и спонтанен, не позволяешь планам легко тебя ограничивать и хорошо реагируешь на изменения.`,
				"medium": `При Добросовестности средний показатель означает, что ты можешь как действовать методично, так и спонтанно реагировать на новые ситуации.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Эмоциональной стабильности высокий показатель означает, что ты сохраняешь спокойствие под давлением, не легко выбиваешься из колеи неудачами и эмоционально уравновешен.`,
				"low":    `При Эмоциональной стабильности низкий показатель означает, что ты интенсивно переживаешь чувства, чутко реагируешь на окружение и обладаешь высоким эмоциональным восприятием.`,
				"medium": `При Эмоциональной стабильности средний показатель означает, что ты можешь как интенсивно переживать чувства, так и спокойно с ними справляться.`,
			},
			domain.TraitOpenness: {
				"high":   `При Открытости опыту высокий показатель означает, что ты любопытен к новому опыту, любишь мыслить креативно и увлекаешься абстрактными идеями.`,
				"low":    `При Открытости опыту низкий показатель означает, что ты практически ориентирован, ценишь проверенные методы и комфортно чувствуешь себя в знакомых структурах.`,
				"medium": `При Открытости опыту средний показатель означает, что ты как ценишь новый опыт, так и умеешь следовать проверенным путям.`,
			},
		}

	case "pl":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Dla Ekstrawersji wysoki wynik oznacza, że czerpiesz energię z kontaktu z innymi, lubisz być w centrum uwagi i czujesz się komfortowo w grupach.`,
				"low":    `Dla Ekstrawersji niski wynik oznacza, że czerpiesz energię raczej ze spokojnych chwil sam na sam i wolisz głębokie rozmowy z kilkoma osobami od dużych grup.`,
				"medium": `Dla Ekstrawersji średni wynik oznacza, że możesz elastycznie przełączać się między momentami towarzyskimi a czasem dla siebie – w zależności od sytuacji.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Dla Ugodowości wysoki wynik oznacza, że harmonia jest dla ciebie ważna, lubisz współpracować i często stawiasz potrzeby innych na pierwszym miejscu.`,
				"low":    `Dla Ugodowości niski wynik oznacza, że wyraźnie wyrażasz własną opinię, nie unikasz konfliktów i nie poddajesz się łatwo wpływom innych.`,
				"medium": `Dla Ugodowości średni wynik oznacza, że możesz być zarówno kooperatywny, jak i wyraźnie reprezentować własne interesy – w zależności od sytuacji.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Dla Sumienności wysoki wynik oznacza, że pracujesz w sposób uporządkowany i zorganizowany, niezawodnie wykonujesz zadania i konsekwentnie dążysz do długoterminowych celów.`,
				"low":    `Dla Sumienności niski wynik oznacza, że jesteś elastyczny i spontaniczny, nie pozwalasz planom łatwo cię ograniczać i dobrze reagujesz na zmiany.`,
				"medium": `Dla Sumienności średni wynik oznacza, że możesz zarówno postępować metodycznie, jak i spontanicznie reagować na nowe sytuacje.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Dla Stabilności emocjonalnej wysoki wynik oznacza, że pozostajesz spokojny pod presją, nie dajesz się łatwo wytrącić z równowagi niepowodzeniami i jesteś emocjonalnie zrównoważony.`,
				"low":    `Dla Stabilności emocjonalnej niski wynik oznacza, że intensywnie przeżywasz uczucia, wrażliwie reagujesz na otoczenie i masz wysoką świadomość emocjonalną.`,
				"medium": `Dla Stabilności emocjonalnej średni wynik oznacza, że możesz zarówno intensywnie przeżywać uczucia, jak i spokojnie sobie z nimi radzić.`,
			},
			domain.TraitOpenness: {
				"high":   `Dla Otwartości na doświadczenia wysoki wynik oznacza, że jesteś ciekawy nowych doświadczeń, lubisz myśleć kreatywnie i ekscytujesz się abstrakcyjnymi ideami.`,
				"low":    `Dla Otwartości na doświadczenia niski wynik oznacza, że jesteś praktycznie zorientowany, cenisz sprawdzone metody i czujesz się komfortowo w znanych strukturach.`,
				"medium": `Dla Otwartości na doświadczenia średni wynik oznacza, że zarówno cenisz nowe doświadczenia, jak i wiesz, jak podążać sprawdzonymi ścieżkami.`,
			},
		}

	case "ro":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Pentru Extraversie, un scor mare înseamnă că îți tragi energia din contactul cu alții, îți place să fii în centrul atenției și te simți confortabil în grupuri.`,
				"low":    `Pentru Extraversie, un scor mic înseamnă că îți tragi energia mai degrabă din momentele liniștite petrecute singur și preferi conversațiile profunde cu câțiva oameni în locul grupurilor mari.`,
				"medium": `Pentru Extraversie, un scor mediu înseamnă că poți comuta flexibil între momentele sociale și timpul pentru tine – în funcție de ceea ce necesită situația.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Pentru Agreabilitate, un scor mare înseamnă că armonia este importantă pentru tine, îți place să cooperezi și adesea pui nevoile altora pe primul loc.`,
				"low":    `Pentru Agreabilitate, un scor mic înseamnă că îți exprimi clar propria opinie, nu te ferești de conflicte și nu ești ușor influențat de alții.`,
				"medium": `Pentru Agreabilitate, un scor mediu înseamnă că poți fi atât cooperant, cât și să îți reprezinți clar propriile interese – în funcție de situație.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Pentru Conștiinciozitate, un scor mare înseamnă că lucrezi într-un mod structurat și organizat, finalizezi sarcinile în mod fiabil și urmărești consecvent obiectivele pe termen lung.`,
				"low":    `Pentru Conștiinciozitate, un scor mic înseamnă că ești flexibil și spontan, nu lași planurile să te constrângă ușor și poți răspunde bine la schimbări.`,
				"medium": `Pentru Conștiinciozitate, un scor mediu înseamnă că poți atât să procedezi metodic, cât și să reacționezi spontan la situații noi.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Pentru Stabilitatea emoțională, un scor mare înseamnă că rămâi calm sub presiune, nu ești ușor descurajat de eșecuri și ești echilibrat emoțional.`,
				"low":    `Pentru Stabilitatea emoțională, un scor mic înseamnă că trăiești sentimentele intens, reacționezi sensibil la mediul tău și ai o conștientizare emoțională ridicată.`,
				"medium": `Pentru Stabilitatea emoțională, un scor mediu înseamnă că poți atât să trăiești sentimentele intens, cât și să le gestionezi calm.`,
			},
			domain.TraitOpenness: {
				"high":   `Pentru Deschidere, un scor mare înseamnă că ești curios despre experiențe noi, îți place să gândești creativ și ești entuziasmat de ideile abstracte.`,
				"low":    `Pentru Deschidere, un scor mic înseamnă că ești orientat practic, apreciezi metodele dovedite și te simți confortabil în structuri familiare.`,
				"medium": `Pentru Deschidere, un scor mediu înseamnă că atât apreciezi experiențele noi, cât și știi cum să urmezi căile dovedite.`,
			},
		}

	case "it":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Per l'Estroversione, un punteggio alto significa che trai energia dal contatto con gli altri, ti piace essere al centro dell'attenzione e ti senti a tuo agio nei gruppi.`,
				"low":    `Per l'Estroversione, un punteggio basso significa che trai la tua energia più dai momenti tranquilli da solo e preferisci conversazioni profonde con poche persone rispetto ai grandi gruppi.`,
				"medium": `Per l'Estroversione, un punteggio medio significa che puoi passare in modo flessibile tra momenti sociali e tempo per te stesso – a seconda di ciò che la situazione richiede.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Per l'Amicalità, un punteggio alto significa che l'armonia è importante per te, ti piace cooperare e spesso metti i bisogni degli altri al primo posto.`,
				"low":    `Per l'Amicalità, un punteggio basso significa che esprimi chiaramente la tua opinione, non rifuggi dai conflitti e non sei facilmente influenzato dagli altri.`,
				"medium": `Per l'Amicalità, un punteggio medio significa che puoi essere sia cooperativo sia rappresentare chiaramente i tuoi interessi – a seconda della situazione.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Per la Coscienziosità, un punteggio alto significa che lavori in modo strutturato e organizzato, completi i compiti in modo affidabile e persegui costantemente obiettivi a lungo termine.`,
				"low":    `Per la Coscienziosità, un punteggio basso significa che sei flessibile e spontaneo, non ti lasci facilmente vincolare dai piani e puoi rispondere bene ai cambiamenti.`,
				"medium": `Per la Coscienziosità, un punteggio medio significa che puoi sia procedere metodicamente sia reagire spontaneamente a nuove situazioni.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Per la Stabilità emotiva, un punteggio alto significa che rimani calmo sotto pressione, non ti lasci facilmente scoraggiare dalle battute d'arresto e sei emotivamente equilibrato.`,
				"low":    `Per la Stabilità emotiva, un punteggio basso significa che vivi i sentimenti intensamente, reagisci sensibilmente al tuo ambiente e hai un'alta consapevolezza emotiva.`,
				"medium": `Per la Stabilità emotiva, un punteggio medio significa che puoi sia vivere i sentimenti intensamente sia gestirli con calma.`,
			},
			domain.TraitOpenness: {
				"high":   `Per l'Apertura mentale, un punteggio alto significa che sei curioso di nuove esperienze, ti piace pensare creativamente e sei entusiasta delle idee astratte.`,
				"low":    `Per l'Apertura mentale, un punteggio basso significa che sei orientato praticamente, apprezzi i metodi collaudati e ti senti a tuo agio nelle strutture familiari.`,
				"medium": `Per l'Apertura mentale, un punteggio medio significa che apprezzi sia le nuove esperienze sia sai come seguire percorsi collaudati.`,
			},
		}

	case "uk":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Екстраверсії високий показник означає, що ти черпаєш енергію зі спілкування з іншими, любиш бути в центрі уваги і комфортно почуваєшся в групах.`,
				"low":    `При Екстраверсії низький показник означає, що ти черпаєш енергію радше зі спокійних моментів наодинці і надаєш перевагу глибоким розмовам з кількома людьми, а не великим групам.`,
				"medium": `При Екстраверсії середній показник означає, що ти можеш гнучко перемикатися між соціальними моментами і часом для себе – залежно від ситуації.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Доброзичливості високий показник означає, що гармонія для тебе важлива, ти любиш співпрацювати і часто ставиш потреби інших на перше місце.`,
				"low":    `При Доброзичливості низький показник означає, що ти чітко висловлюєш власну думку, не уникаєш конфліктів і не легко піддаєшся впливу інших.`,
				"medium": `При Доброзичливості середній показник означає, що ти можеш бути як кооперативним, так і чітко відстоювати власні інтереси – залежно від ситуації.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Сумлінності високий показник означає, що ти працюєш структуровано й організовано, надійно виконуєш завдання і послідовно переслідуєш довгострокові цілі.`,
				"low":    `При Сумлінності низький показник означає, що ти гнучкий і спонтанний, не дозволяєш планам легко тебе обмежувати і добре реагуєш на зміни.`,
				"medium": `При Сумлінності середній показник означає, що ти можеш як діяти методично, так і спонтанно реагувати на нові ситуації.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Емоційній стабільності високий показник означає, що ти зберігаєш спокій під тиском, не легко вибиваєшся з колії невдачами і емоційно врівноважений.`,
				"low":    `При Емоційній стабільності низький показник означає, що ти інтенсивно переживаєш почуття, чутливо реагуєш на оточення і маєш високе емоційне сприйняття.`,
				"medium": `При Емоційній стабільності середній показник означає, що ти можеш як інтенсивно переживати почуття, так і спокійно з ними справлятися.`,
			},
			domain.TraitOpenness: {
				"high":   `При Відкритості досвіду високий показник означає, що ти цікавишся новим досвідом, любиш мислити креативно і захоплюєшся абстрактними ідеями.`,
				"low":    `При Відкритості досвіду низький показник означає, що ти практично орієнтований, цінуєш перевірені методи і комфортно почуваєшся в знайомих структурах.`,
				"medium": `При Відкритості досвіду середній показник означає, що ти як цінуєш новий досвід, так і вмієш слідувати перевіреним шляхам.`,
			},
		}

	case "bg":
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `При Екстраверсия висок резултат означава, че черпиш енергия от контакта с другите, обичаш да си в центъра на вниманието и се чувстваш комфортно в групи.`,
				"low":    `При Екстраверсия нисък резултат означава, че черпиш енергия повече от спокойни моменти насаме и предпочиташ дълбоки разговори с малко хора пред големи групи.`,
				"medium": `При Екстраверсия среден резултат означава, че можеш гъвкаво да превключваш между социални моменти и време за себе си – в зависимост от ситуацията.`,
			},
			domain.TraitAgreeableness: {
				"high":   `При Сговорчивост висок резултат означава, че хармонията е важна за теб, обичаш да сътрудничиш и често поставяш нуждите на другите на първо място.`,
				"low":    `При Сговорчивост нисък резултат означава, че ясно изразяваш собственото си мнение, не се страхуваш от конфликти и не се поддаваш лесно на влиянието на другите.`,
				"medium": `При Сговорчивост среден резултат означава, че можеш да бъдеш както кооперативен, така и ясно да представяш собствените си интереси – в зависимост от ситуацията.`,
			},
			domain.TraitConscientiousness: {
				"high":   `При Съвестност висок резултат означава, че работиш по структуриран и организиран начин, надеждно изпълняваш задачи и последователно преследваш дългосрочни цели.`,
				"low":    `При Съвестност нисък резултат означава, че си гъвкав и спонтанен, не позволяваш на плановете лесно да те ограничават и можеш да реагираш добре на промени.`,
				"medium": `При Съвестност среден резултат означава, че можеш както да действаш методично, така и спонтанно да реагираш на нови ситуации.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `При Емоционална стабилност висок резултат означава, че оставаш спокоен под напрежение, не се разколебаваш лесно от неуспехи и си емоционално балансиран.`,
				"low":    `При Емоционална стабилност нисък резултат означава, че изпитваш чувствата интензивно, реагираш чувствително към заобикалящата те среда и имаш високо емоционално осъзнаване.`,
				"medium": `При Емоционална стабилност среден резултат означава, че можеш както да изпитваш чувствата интензивно, така и спокойно да се справяш с тях.`,
			},
			domain.TraitOpenness: {
				"high":   `При Отвореност висок резултат означава, че си любопитен към нови преживявания, обичаш да мислиш креативно и си ентусиазиран от абстрактни идеи.`,
				"low":    `При Отвореност нисък резултат означава, че си практически ориентиран, цениш доказани методи и се чувстваш комфортно в познати структури.`,
				"medium": `При Отвореност среден резултат означава, че както цениш нови преживявания, така и знаеш как да следваш доказани пътища.`,
			},
		}

	default: // German (default)
		return map[domain.Trait]map[string]string{
			domain.TraitExtraversion: {
				"high":   `Bei Extraversion bedeutet ein hoher Wert, dass du Energie aus dem Kontakt mit anderen Menschen ziehst, gerne im Mittelpunkt stehst und dich in Gruppen wohlfühlst.`,
				"low":    `Bei Extraversion bedeutet ein niedriger Wert, dass du deine Energie eher aus ruhigen Momenten für dich allein schöpfst und tiefe Gespräche mit wenigen Menschen großen Gruppen vorziehst.`,
				"medium": `Bei Extraversion bedeutet ein mittlerer Wert, dass du flexibel zwischen geselligen Momenten und Zeit für dich wechseln kannst – je nachdem, was die Situation erfordert.`,
			},
			domain.TraitAgreeableness: {
				"high":   `Bei Verträglichkeit bedeutet ein hoher Wert, dass dir Harmonie wichtig ist, du gerne kooperierst und die Bedürfnisse anderer oft an erste Stelle setzt.`,
				"low":    `Bei Verträglichkeit bedeutet ein niedriger Wert, dass du deine eigene Meinung klar vertrittst, Konflikte nicht scheust und dich nicht so leicht von anderen beeinflussen lässt.`,
				"medium": `Bei Verträglichkeit bedeutet ein mittlerer Wert, dass du sowohl kooperativ sein kannst als auch deine eigenen Interessen klar vertreten kannst – situationsabhängig.`,
			},
			domain.TraitConscientiousness: {
				"high":   `Bei Gewissenhaftigkeit bedeutet ein hoher Wert, dass du strukturiert und organisiert vorgehst, Aufgaben zuverlässig erledigst und langfristige Ziele konsequent verfolgst.`,
				"low":    `Bei Gewissenhaftigkeit bedeutet ein niedriger Wert, dass du flexibel und spontan bist, dich nicht so leicht von Plänen einengen lässt und gut auf Veränderungen reagieren kannst.`,
				"medium": `Bei Gewissenhaftigkeit bedeutet ein mittlerer Wert, dass du sowohl planvoll vorgehen als auch spontan auf neue Situationen reagieren kannst.`,
			},
			domain.TraitEmotionalStability: {
				"high":   `Bei Emotionaler Stabilität bedeutet ein hoher Wert, dass du auch unter Druck gelassen bleibst, dich von Rückschlägen nicht so leicht aus der Bahn werfen lässt und emotional ausgeglichen bist.`,
				"low":    `Bei Emotionaler Stabilität bedeutet ein niedriger Wert, dass du Gefühle intensiv erlebst, feinfühlig auf deine Umgebung reagierst und eine hohe emotionale Wahrnehmung hast.`,
				"medium": `Bei Emotionaler Stabilität bedeutet ein mittlerer Wert, dass du Gefühle sowohl intensiv erleben als auch gelassen damit umgehen kannst.`,
			},
			domain.TraitOpenness: {
				"high":   `Bei Offenheit bedeutet ein hoher Wert, dass du neugierig auf neue Erfahrungen bist, gerne kreativ denkst und abstrakte Ideen dich begeistern.`,
				"low":    `Bei Offenheit bedeutet ein niedriger Wert, dass du praktisch orientiert bist, bewährte Methoden schätzt und dich in vertrauten Strukturen wohlfühlst.`,
				"medium": `Bei Offenheit bedeutet ein mittlerer Wert, dass du sowohl neue Erfahrungen schätzt als auch bewährte Wege zu gehen weißt.`,
			},
		}
	}
}
