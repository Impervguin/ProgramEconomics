export const PARAM_LEVELS = {
  NoInfluence: 'NoInfluence',
  RandomInfluence: 'RandomInfluence',
  SmallInfluence: 'SmallInfluence',
  MediumInfluence: 'MediumInfluence',
  ImportantInfluence: 'ImportantInfluence',
  MainInfluence: 'MainInfluence',
}

export const PARAM_LEVEL_LABELS = {
  NoInfluence: 'Нет влияния',
  RandomInfluence: 'Случайное влияние',
  SmallInfluence: 'Небольшое влияние',
  MediumInfluence: 'Среднее влияние',
  ImportantInfluence: 'Важное влияние',
  MainInfluence: 'Главное влияние',
}

export const LANGUAGES = {
  Assembly: 'Assembly',
  C: 'C',
  Cobol: 'Cobol',
  Fortran: 'Fortran',
  Pascal: 'Pascal',
  CPlusPlus: 'C++',
  CSharp: 'C#',
  Java: 'Java',
  Ada95: 'Ada95',
  VisualBasic: 'VisualBasic',
  VisualCpp: 'VisualCpp',
  Delphi: 'Delphi',
  Perl: 'Perl',
  Prolog: 'Prolog',
  SQL: 'SQL',
  Lisp: 'Lisp',
}

export const ELEMENTS = {
  externalInput: 'Внешний ввод',
  externalOutput: 'Внешний вывод',
  externalQuery: 'Внешний запрос',
  internalLogicalFile: 'Внутренний логический файл',
  externalInterfaceFile: 'Внешний интерфейсный файл',
}

export const ELEMENT_LEVELS = {
  Low: 'Low',
  Normal: 'Normal',
  High: 'High',
}

export const ELEMENT_LEVEL_LABELS = {
  Low: 'Низкая',
  Normal: 'Средняя',
  High: 'Высокая',
}

export const TEAM_PARAM_LEVELS = {
  Team: {
    SeverelyHindered: 'SeverelyHindered',
    SomewhatHindered: 'SomewhatHindered',
    SomeCohesion: 'SomeCohesion',
    IncreasedCohesion: 'IncreasedCohesion',
    HighCohesion: 'HighCohesion',
    UnifiedWhole: 'UnifiedWhole',
  },
  Pmat: {
    CMMLevel1: 'CMMLevel1',
    CMMLevel1Plus: 'CMMLevel1Plus',
    CMMLevel2: 'CMMLevel2',
    CMMLevel3: 'CMMLevel3',
    CMMLevel4: 'CMMLevel4',
    CMMLevel5: 'CMMLevel5',
  },
  Prec: {
    NewChaoticProject: 'NewChaoticProject',
    NewProject: 'NewProject',
    SomeExperienceProject: 'SomeExperienceProject',
    CommonExperienceProject: 'CommonExperienceProject',
    ExperiencedProject: 'ExperiencedProject',
    FullyKnownProject: 'FullyKnownProject',
  },
  Resl: {
    Little20: 'Little20',
    Some40: 'Some40',
    Frequent60: 'Frequent60',
    Overall75: 'Overall75',
    AlmostFull90: 'AlmostFull90',
    Full100: 'Full100',
  },
  Flex: {
    RigidProcess: 'RigidProcess',
    RandomRelaxations: 'RandomRelaxations',
    SomeRelaxations: 'SomeRelaxations',
    MostlyAgreedProcess: 'MostlyAgreedProcess',
    SomeAgreement: 'SomeAgreement',
    GeneralGoalsOnly: 'GeneralGoalsOnly',
  },
}

export const TEAM_PARAM_LABELS = {
  Team: {
    SeverelyHindered: 'Сильно затрудненное взаимодействие',
    SomewhatHindered: 'Несколько затрудненное взаимодействие',
    SomeCohesion: 'Некоторая согласованность',
    IncreasedCohesion: 'Повышенная согласованность',
    HighCohesion: 'Высокая согласованность',
    UnifiedWhole: 'Взаимодействие как в едином целом',
  },
  Pmat: {
    CMMLevel1: 'Уровень 1 СММ',
    CMMLevel1Plus: 'Уровень 1+ СММ',
    CMMLevel2: 'Уровень 2 СММ',
    CMMLevel3: 'Уровень 3 СММ',
    CMMLevel4: 'Уровень 4 СММ',
    CMMLevel5: 'Уровень 5 СММ',
  },
  Prec: {
    NewChaoticProject: 'Полное отсутствие прецедентов, полностью непредсказуемый проект',
    NewProject: 'Почти полное отсутствие прецедентов, в значительной мере непредсказуемый проект',
    SomeExperienceProject: 'Наличие некоторого количества прецедентов',
    CommonExperienceProject: 'Общее знакомство с проектом',
    ExperiencedProject: 'Значительное знакомство с проектом',
    FullyKnownProject: 'Полное знакомство с проектом',
  },
  Resl: {
    Little20: 'Малое (20 %)',
    Some40: 'Некоторое (40 %)',
    Frequent60: 'Частое (60 %)',
    Overall75: 'В целом (75 %)',
    AlmostFull90: 'Почти полное (90 %)',
    Full100: 'Полное (100%)',
  },
  Flex: {
    RigidProcess: 'Точный, строгий процесс разработки',
    RandomRelaxations: 'Случайные послабления в процессе',
    SomeRelaxations: 'Некоторые послабления в процессе',
    MostlyAgreedProcess: 'Большей частью согласованный процесс',
    SomeAgreement: 'Некоторое согласование процесса',
    GeneralGoalsOnly: 'Заказчик определил только общие цели',
  },
}

export const PRODUCTIVITY_LEVELS = {
  VeryLowProductivity: 'VeryLowProductivity',
  LowProductivity: 'LowProductivity',
  MediumProductivity: 'MediumProductivity',
  HighProductivity: 'HighProductivity',
  VeryHighProductivity: 'VeryHighProductivity',
}

export const PRODUCTIVITY_LABELS = {
  VeryLowProductivity: 'Очень низкая',
  LowProductivity: 'Низкая',
  MediumProductivity: 'Средняя',
  HighProductivity: 'Высокая',
  VeryHighProductivity: 'Очень высокая',
}

export const ARCHITECTURE_ATTRIBUTES = {
  PERS: 'PERS',
  RCPX: 'RCPX',
  RUSE: 'RUSE',
  PDIF: 'PDIF',
  PREX: 'PREX',
  FCIL: 'FCIL',
  SCED: 'SCED',
}

export const ATTRIBUTE_LEVELS = {
  ExtraLow: 'ExtraLow',
  VeryLow: 'VeryLow',
  Low: 'Low',
  Nominal: 'Nominal',
  High: 'High',
  VeryHigh: 'VeryHigh',
  ExtraHigh: 'ExtraHigh',
}

export const ATTRIBUTE_LABELS = {
  ExtraLow: 'Очень низкий',
  VeryLow: 'Низкий',
  Low: 'Низкий',
  Nominal: 'Номинальный',
  High: 'Высокий',
  VeryHigh: 'Очень высокий',
  ExtraHigh: 'Сверхвысокий',
}

export const VIEW_REPORT_LEVELS = {
  Simple: 'Simple',
  Medium: 'Medium',
  Complex: 'Complex',
}

export const VIEW_REPORT_LABELS = {
  Simple: 'Простой',
  Medium: 'Средний',
  Complex: 'Сложный',
}