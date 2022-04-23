# -*- coding: utf-8 -*-

import os
import re
import sys
import keyboard
import signal
import math
import rsa
import time
import datetime
import json
import subprocess
import logging
import traceback
import fileinput
import itertools
import asyncio
import aiohttp
import requests
import urllib3
import urllib
import urllib.request
import hmac, hashlib
import ssl
import sqlite3
import telebot
import threading
import async_timeout
import twisted.internet
from os import system, name
from sys import exc_info
from traceback import extract_tb
from colorama import init
from termcolor import cprint, colored
from threading import Thread
from binance.websockets import BinanceSocketManager
from binance.client import Client
from twisted.internet import reactor
from twisted.protocols.basic import LineReceiver
from twisted.internet.protocol import Factory
from requests.packages.urllib3.exceptions import InsecureRequestWarning
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)
init()

class Initialization(): 
    """Инициализация начальных переменных"""

    def __init__(self):
        self.version_bot = self.last_version_bot = '0.312' # Объявляем версии бота
        self.new_version = '' # Объявляем версии бота
        self.color_version = 'green' # Цвет текущей версии бота
        self.black_asset = ['UP', 'DOWN', 'BULL', 'BEAR'] # Чёрный список активов
        self.system_key_global = [ # Не торговые ключи в базе данных
            'symbols',
            'trade_info',
            'api_key',
            'trade_params',
            'trade_params_list',
            'white_list',
            'trade_pairs',
            'trailing_orders',
            'daily_profit',
            'bnb_burn',
            'average_percent']
        self.trade_info_global = { # Торговая статистика
            'sell_filled_orders': '0',
            'sell_open_orders': '0'}
        self.api_key_global = { # API ключи
            'api': '',
            'secret': '',
            'referral': '',
            'bep20': '',
            'tg_notification': 'False',
            'tg_token': '0:A-s',
            'tg_name': '@'}
        self.trade_params_global = { # Торговые параметры бота
            'name_list': 'native',
            'min_bnb': '0.01',
            'min_balance': '25',
            'min_order': '1.1',
            'min_price': '0.005',
            'daily_percent': '-5',
            'sell_up': '1.75',
            'buy_down': '-5',
            'max_trade_pairs': '20',
            'auto_trade_pairs': True,
            'delta_percent': True,
            'num_aver': True,
            'step_aver': '1.15',
            'max_aver': '4',
            'quantity_aver': '2',
            'trailing_stop': False,
            'trailing_percent': '0.35',
            'trailing_part': '10',
            'trailing_price': '0.15',
            'user_order': True,
            'fiat_currencies': 'RUB',
            'quote_asset': 'USDT BTC',
            'double_asset': False}

    def github(self):
        """Сверяем версию бота и актуальную на GitHub"""
        try:
            response = requests.get("https://api.github.com/repos/test") # Обращаемся к GitHub
            var.last_version_bot = response.json()['tag_name'] # Получаем последнюю версию бота по тэгу
            if float(var.last_version_bot) > float(var.version_bot): # Если текущая версия бота меньше, чем актуальная версия на GitHub, то
                var.new_version = colored('\nДоступна новая версия: ', 'cyan') + colored(var.last_version_bot, 'green') # Оповещаем пользователя в консоли о новой версии
            self.color_version = 'green' if float(var.last_version_bot) <= float(var.version_bot) else 'red'
        except Exception as e:
            logging.error('init.github():\nresponse: {}\nexcept: {}\n'.format(response, str(e)))

class MathFunc():
    """Операции с числами"""

    def get_count(self, number):
        """Сколько знаков после запятой?

        Args:
            number (float): получаем число

        Returns:
            int: получаем количество знаков после запятой
        """
        self.str_number = str(number)
        if '.' not in self.str_number:
            return 0
        return len(number[number.index('.') + 1:])

    def number_round(self, number):
        """Форматирование float данных

        Args:
            number (float): получаем число

        Returns:
            str: возвращаем число в виде строки с нужным количеством знаков после запятой
        """

        self.str_number = str(number)
        if '.0' in self.str_number:
            self.str_number = self.str_number.rstrip('0')
            self.str_number = self.str_number.rstrip('.')
        elif self.str_number.rstrip('0') and '.' in self.str_number:
            self.str_number = self.str_number.rstrip('0')
        else:
            pass
        return self.str_number

class DataBase(): 
    """Действия с БД"""

    def __init__(self):
        self.connect = sqlite3.connect("xbot_db.db", check_same_thread = False) # или :memory: чтобы сохранить в RAM
        self.cursor = self.connect.cursor()

    def check_tables(self):
        """Проверка БД на наличие таблиц

        Migrations for DB

        Если база только создана и в ней нет API ключей, то задаем вопросы в консоль и заполняем переменные.
        """

        db.cursor.execute("""
            CREATE TABLE if not EXISTS symbols(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                pair TEXT UNIQUE,
                baseAsset TEXT,
                quoteAsset TEXT,
                stepSize TEXT,
                tickSize TEXT,
                minNotional TEXT,
                priceChangePercent TEXT,
                bidPrice TEXT,
                askPrice TEXT,
                averagePrice TEXT,
                buyPrice TEXT,
                sellPrice TEXT,
                trailingPrice TEXT,
                allQuantity TEXT,
                freeQuantity TEXT,
                lockQuantity TEXT,
                orderId TEXT,
                profit TEXT,
                totalQuote TEXT,
                stepAveraging TEXT,
                numAveraging TEXT,
                statusOrder TEXT)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS trade_info(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                sell_filled_orders TEXT,
                sell_open_orders TEXT)""")
        for self._key, self._value in var.trade_info_global.items():
            self._checked = db.read('trade_info')
            self._checked = self._checked[0] if len(self._checked) > 0 else self._checked
            if self._key not in self._checked or self._checked[self._key] == None:
                db.write('insert', 'trade_info', self._key, self._value) if len(self._checked) == 0 else db.write('update', 'trade_info', self._key, self._value)

        db.cursor.execute("""
            CREATE TABLE if not EXISTS api_key(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                api TEXT,
                secret TEXT,
                referral TEXT,
                tg_notification BOOLEAN,
                tg_token TEXT,
                tg_name TEXT)""")

        for self._key, self._value in var.api_key_global.items():
            self._checked = db.read('api_key')
            self._checked = self._checked[0] if len(self._checked) > 0 else self._checked
            if self._key not in self._checked or self._checked[self._key] == None:
                if self._key == 'api':
                    cprint('\033[H\033[J\033[K' + colored('Введите API ключ от Binance:\033[K', 'cyan'))
                    self._value = input('')
                elif self._key == 'secret':
                    cprint('\033[H\033[J\033[K' + colored('Введите Secret ключ от Binance:\033[K', 'cyan'))
                    self._value = input('')
                elif self._key == 'referral':
                    cprint('\033[H\033[J\033[K' + colored('Введите Ваш ID пользователя на Binance (состоит только из цифр):\033[K', 'cyan'))
                    while True:
                        try:
                            self._value = str(abs(int(input(''))))
                            break
                        except:
                            print('\033[A\033[K\033[H')
                    print('\033[H\033[J\033[K')
                elif self._key == 'bep20':
                    while True:
                        cprint('\033[H\033[J\033[K' + colored('Введите депозитный адрес BUSD в сети BEP20 (Binance Smart Chain) человека, который вас пригласил, или оставьте поле пустым:\033[K', 'cyan'))
                        self._value = list(map(str, input('').split(' ')))
                        if len(self._value) == 1:
                            for _ in self._value:
                                print(_[:2])
                                self._value = 'Отсутствует' if len(_) == 0 else _ if len(_) == 42 and _[:2] == '0x' else ''
                                true_wallet = True if self._value != '' else False
                            if true_wallet == True:
                                break
                db.write('insert', 'api_key', self._key, self._value) if len(self._checked) == 0 else db.write('update', 'api_key', self._key, self._value)

        db.cursor.execute("""
            CREATE TABLE if not EXISTS trade_params(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                name_list TEXT,
                min_bnb TEXT,
                min_balance TEXT,
                min_order TEXT,
                min_price TEXT,
                daily_percent TEXT,
                sell_up TEXT,
                buy_down TEXT,
                max_trade_pairs TEXT,
                auto_trade_pairs BOOLEAN,
                delta_percent BOOLEAN,
                num_aver BOOLEAN,
                step_aver TEXT,
                max_aver TEXT,
                quantity_aver TEXT,
                trailing_stop BOOLEAN,
                trailing_percent TEXT,
                trailing_part TEXT,
                trailing_price TEXT,
                user_order BOOLEAN,
                fiat_currencies TEXT,
                quote_asset TEXT,
                double_asset BOOLEAN)""")
        for self._key, self._value in var.trade_params_global.items():
            self._checked = db.read('trade_params')
            self._checked = self._checked[0] if len(self._checked) > 0 else self._checked
            if self._key not in self._checked or self._checked[self._key] == None:
                db.write('insert', 'trade_params', self._key, self._value) if len(self._checked) == 0 else db.write('update', 'trade_params', self._key, self._value)

        db.cursor.execute("""
            CREATE TABLE if not EXISTS trade_params_list(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                name_list TEXT,
                min_bnb TEXT,
                min_balance TEXT,
                min_order TEXT,
                min_price TEXT,
                daily_percent TEXT,
                sell_up TEXT,
                buy_down TEXT,
                max_trade_pairs TEXT,
                auto_trade_pairs BOOLEAN,
                delta_percent BOOLEAN,
                num_aver BOOLEAN,
                step_aver TEXT,
                max_aver TEXT,
                quantity_aver TEXT,
                trailing_stop BOOLEAN,
                trailing_percent TEXT,
                trailing_part TEXT,
                trailing_price TEXT,
                user_order BOOLEAN,
                fiat_currencies TEXT,
                quote_asset TEXT,
                double_asset BOOLEAN)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS white_list(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                pair TEXT)""")
        self.dev_white_list = [ # Первоначальный белый список (высокий риск)
                'ADA',
                'ADX',
                'AGI',
                'AION',
                'ALGO',
                'ARDR',
                'ARPA',
                'ATOM',
                'BCH',
                'BLZ',
                'BNT',
                'COTI',
                'CVC',
                'DASH',
                'DATA',
                'DCR',
                'ELF',
                'ENJ',
                'EOS',
                'ETH',
                'GXS',
                'ICX',
                'IOTA',
                'IRIS',
                'KMD',
                'LINK',
                'LOOM',
                'LSK',
                'LTC',
                'LTO',
                'MANA',
                'NANO',
                'NEO',
                'NULS',
                'OAX',
                'OGN',
                'OMG',
                'ONT',
                'PERL',
                'POA',
                'POLY',
                'QLC',
                'QSP',
                'QTUM',
                'REN',
                'REP',
                'RLC',
                'RVN',
                'SNT',
                'STEEM',
                'STORJ',
                'SXP',
                'THETA',
                'WAN',
                'WAVES',
                'WRX',
                'XEM',
                'XLM',
                'XMR',
                'XTZ',
                'ZEN',
                'ZIL',
                'ZRX']
        if len(db.read('sqlite_sequence', condition = "WHERE name = 'white_list'", keys = ['name'])) == 0:
            self._key = 'pair'
            for self._value in self.dev_white_list:
                self._checked = db.read('white_list', condition = "WHERE {} = '{}'".format(self._key, self._value))
                self._checked = self._checked[0][self._key] if len(self._checked) > 0 else self._checked
                if self._value not in self._checked:
                    db.write('insert', 'white_list', self._key, self._value)

        db.cursor.execute("""
            CREATE TABLE if not EXISTS trade_pairs(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                pair TEXT,
                baseAsset TEXT,
                quoteAsset TEXT)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS trailing_orders(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                pair TEXT,
                p TEXT,
                q TEXT)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS daily_profit(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                day TEXT,
                quote TEXT,
                profit TEXT)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS bnb_burn(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                day TEXT,
                pair TEXT,
                size TEXT,
                comission TEXT)""")

        db.cursor.execute("""
            CREATE TABLE if not EXISTS average_percent(
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                day TEXT,
                percent TEXT)""")
        self.connect.commit()

    def check_items(self):
        """Проверка и сопоставление информации по торгам и статистики"""
        [db.write('delete', 'symbols', 'pair', symbol['pair']) for symbol in db.read('symbols', condition = "WHERE statusOrder LIKE 'NO_ORDER'", keys = ['pair', 'baseAsset', 'quoteAsset']) if symbol['baseAsset'] not in bot.white_list or symbol['quoteAsset'] not in bot.quote_asset_list]

    def check_table_trade_pairs(self):
        """Проверка таблицы trade_pairs"""
        db.cursor.executescript("""
            DELETE FROM 'trade_pairs';
            UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='trade_pairs';""")
        for self._value in main.exchange['symbols']:
            self._checked = db.read('trade_pairs', condition = "WHERE {} = '{}'".format('pair', self._value['symbol']))
            self._checked = self._checked[0]['pair'] if len(self._checked) > 0 else self._checked
            if self._value['symbol'] not in self._checked and self._value['quoteAsset'] in list(set(map(str, db.read('trade_params')[0]['quote_asset'].split(' ')))) and False not in [False for _ in var.black_asset if _ in self._value['baseAsset']] and self._value['status'] == 'TRADING':
                db.write('insert', 'trade_pairs', 'pair', self._value['symbol'], pair = self._value['symbol'], baseAsset = self._value['baseAsset'], quoteAsset = self._value['quoteAsset'])

    def write(self, _action, _table, _key, _value, _dict = None, **kwargs):
        """Пишем в БД"""
        if _table in var.system_key_global:
            try:
                _k = _v = _kv = ""

                if _action == 'insert':
                    for key, value in kwargs.items():
                        _k += "{},".format(key)
                        _v += "'{}',".format(value)
                    _key = "'" + _key + "'" if len(kwargs) == 0 else _k[:-1]
                    _value = "'" + _value + "'" if len(kwargs) == 0 else _v[:-1]
                    #test = str("INSERT INTO {} ({}) values({})".format(_table, _key, _value))
                    db.cursor.execute("INSERT INTO {} ({}) values({})".format(_table, _key, _value))

                elif _action == 'update':
                    try:
                        #test = str("UPDATE {} SET {} = '{}' {}".format(_table, _key, _value, "WHERE {} = '{}'".format([k for k in kwargs][0], [kwargs[v] for v in kwargs][0]) if len(kwargs) != 0 else ''))
                        db.cursor.execute("UPDATE {} SET {} = '{}' {}".format(_table, _key, _value, "WHERE {} = '{}'".format([k for k in kwargs][0], [kwargs[v] for v in kwargs][0]) if len(kwargs) != 0 else ''))
                    except Exception as e:
                        if 'no such column: ' in str(e):
                            e = str(e).replace('no such column: ', '')
                            #test = str("ALTER TABLE {} ADD COLUMN {} {};".format(_table, e, 'TEXT' if type(_value) != bool else 'BOOLEAN'))
                            db.cursor.execute("ALTER TABLE {} ADD COLUMN {} {};".format(_table, e, 'TEXT' if type(_value) != bool else 'BOOLEAN'))
                            db.cursor.execute("UPDATE {} SET {} = '{}' {}".format(_table, e, _value, "WHERE {} = '{}'".format([k for k in kwargs][0], [kwargs[v] for v in kwargs][0]) if len(kwargs) != 0 else ''))

                elif _action == 'updates':
                    for key, value in kwargs.items():
                        _kv += "{} = '{}',".format(key, value)
                    #test = str("UPDATE {} SET {} {}".format(_table, _kv[:-1], "WHERE {} = '{}'".format(_key, _value) if _key != '' and _value != '' else '')))
                    db.cursor.execute("UPDATE {} SET {} {}".format(_table, _kv[:-1], "WHERE {} = '{}'".format(_key, _value) if _key != '' and _value != '' else ''))

                elif _action == 'updates_sp' and _dict is not None:
                    for pair, value in _dict.items():
                        for key in value:
                            _kv += "{} = '{}',".format(key, value[key])
                        #test = db.cursor.execute("UPDATE {} SET {} {}".format(_table, _kv[:-1], "WHERE pair = '{}'".format(pair)))
                        db.cursor.execute("UPDATE {} SET {} {}".format(_table, _kv[:-1], "WHERE pair = '{}'".format(pair)))

                elif _action == 'delete':
                    #test = str("DELETE FROM {} WHERE {} = '{}'".format(_table, _key, _value))
                    db.cursor.execute("DELETE FROM {} WHERE {} = '{}'".format(_table, _key, _value))

                else:
                    return

                db.connect.commit()
            except Exception as e:
                logging.error('db.write():\naction: {}\n_table: {}\n_key: {}\n_value: {}\nexcept: {}\n'.format(_action, _table, _key, _value, str(e)))

    def read(self, table, condition = None, **kwargs):
        """Читаем из БД"""
        try:
            keys = ''
            if len(kwargs.keys()) > 0:
                for value in kwargs['keys']:
                    keys += "{},".format(value)
            request = "SELECT {} FROM {} {}".format(keys[:-1] if len(kwargs.keys()) > 0 else '*', table, condition if condition != None else '')
            db.cursor.execute(request)
            column = [description[0] for description in db.cursor.description]
            result = db.cursor.fetchall()
            request_list = list()
            for _ in result:
                request_dict = dict()
                for value in column:
                    request_dict.update({value: _[column.index(value)]})
                request_list.append(request_dict)
            return request_list
        except:
            return ''

class Main(): # Главное меню

    def __init__(self):
        db.check_tables()
        while True:
            try:
                self.client = Client(db.read('api_key')[0]['api'], db.read('api_key')[0]['secret']) # Присваиваем API ключи клиентскому запросу
                self.exchange = self.client.get_exchange_info() # Получаем информацию по каждой паре (Минимальное количество монет в ордере, цена, количество знаков после точки и т.д.)
                self.tickers = self.client.get_ticker() # Получаем информацию о торгах
                break
            except Exception as e:
                logging.error('main.__init__():\nexcept: {}\n'.format(str(e)))

    def clear(self): # Очистить консоль
        try:
            if name == 'nt': # Win
                os.system('cls')
            else: # Linux
                os.system('clear')
        except:
            pass

    def print_menu(self): # Меню вывода в консоль
        main.clear()
        cprint(colored(
            '\n  *       *          * * *       * * *   * * * * *' +
            '\n    *   *            * ', 'cyan') + colored('*', 'green') + colored('  *     *     *      *    ' +
            '\n      *      * * *   * * * *   *   ', 'cyan') + colored('*', 'red') + colored('   *     *    ' +
            '\n    *   *            * ', 'cyan') + colored('* *', 'green') + colored('  *   *     *      *    ' +
            '\n  *       *          * * * *     * * *       *\n\n', 'cyan') + 
        colored('Версия: ', 'cyan') + colored(var.version_bot, var.color_version) + var.new_version +
        colored('\nGitHub: ', 'cyan') + colored('github.com/test', 'magenta', attrs=['bold']) +
        colored('\nTelegram: ', 'cyan') + colored('https://t.me/test', 'magenta', attrs=['bold']))
        max_orders = 'без ограничений' if float(db.read('trade_params')[0]['max_trade_pairs']) == -1 else 'до ' + db.read('trade_params')[0]['max_trade_pairs']
        max_orders = 'только усреднение открытых' if float(db.read('trade_params')[0]['max_trade_pairs']) == 0 else max_orders
        cprint(
            colored('\nУспешных сделок: ', 'cyan') + colored(int(float(db.read('trade_info')[0]['sell_filled_orders'])), 'yellow') +
            colored('\nОрдеров на продажу: ', 'cyan') + colored(int(float(db.read('trade_info')[0]['sell_open_orders'])), 'yellow') +
            colored('\nОграничение позиций: ', 'cyan') + colored(max_orders, 'yellow') +
            colored('\n\n    ', 'white', 'on_white', attrs=['bold']) + colored(' Главное меню', 'white') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -s ', 'grey', 'on_green') + colored(' Запустить X-Bot', 'green') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -o ', 'grey', 'on_cyan') + colored(' Посмотреть открытые позиции', 'cyan') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -m ', 'grey', 'on_cyan') + colored(' Изменить список монет', 'cyan') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -k ', 'grey', 'on_yellow') + colored(' Изменить API и Telegram настройки', 'yellow') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -p ', 'grey', 'on_yellow') + colored(' Изменить торговые параметры бота', 'yellow') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -h ', 'grey', 'on_red') + colored(' Удалить информацию об ордерах', 'red') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -i ', 'grey', 'on_red') + colored(' Удалить торговую статистику', 'red') +
            colored('\n |__', 'white', attrs=['bold']) + colored(' -e ', 'grey', 'on_red') + colored(' Завершить работу', 'red') +
            '\n')
        main.keys()

    def keys(self): # Ключи команд
        self.key = input('')

        if self.key == '-s': # Запуск бота
            bot.connect()

        elif self.key == '-o': # Посмотреть открытые позиции
            while True:
                main.clear()
                sell_orders = 0
                bot_orders = sorted([pair for pair in db.read('symbols')], key = lambda name: name['pair'])
                cprint(
                    colored('    ', 'white', 'on_cyan', attrs=['bold']) + colored('\n |__', 'cyan') + colored(' -ext ', 'grey', 'on_cyan') + colored(' Выход в меню', 'cyan') +
                    colored('\n |__', 'cyan') + colored(' -clp ', 'grey', 'on_green') + colored(' Закрыть позиции в плюсе\n', 'green') +
                    colored(' |__', 'cyan') + colored(' -sll ', 'grey', 'on_yellow') + colored(' Закрыть выбранные позиции\n', 'yellow') +
                    colored(' |__', 'cyan') + colored(' -cls ', 'grey', 'on_red') + colored(' Закрыть все позиции\n', 'red'))
                for pair in bot_orders:
                    if pair['statusOrder'] == 'SELL_ORDER':
                        cprint(colored('SELL ', 'red') + 'LIMIT ' + colored('[' + pair['allQuantity'] + ']', 'grey', 'on_white') + ' ' + colored(pair['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + colored(pair['sellPrice'] + ' ' + pair['quoteAsset'], 'yellow') + colored(' ({:.2f}% | {} {})'.format((float(pair['sellPrice']) - float(pair['askPrice'])) / (float(pair['askPrice']) / 100), mathematical.number_round('{:.8f}'.format(float(pair['askPrice']) * float(pair['allQuantity']) - float(pair['totalQuote']))), pair['quoteAsset']), 'white') + '\033[K')
                        sell_orders += 1
                    elif pair['statusOrder'] == 'TRAILING_STOP_ORDER':
                        cprint(colored('SELL ', 'magenta') + 'TRAILING ' + colored('[' + pair['freeQuantity'] + ']', 'grey', 'on_white') + ' ' + colored(pair['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + colored(pair['sellPrice'] + ' ' + pair['quoteAsset'], 'yellow') + colored(' ({:.2f}% | {} {})'.format((float(pair['sellPrice']) - float(pair['askPrice'])) / (float(pair['askPrice']) / 100), mathematical.number_round('{:.8f}'.format(float(pair['askPrice']) * float(pair['freeQuantity']) - float(pair['totalQuote']))), pair['quoteAsset']), 'white') + '\033[K')
                        sell_orders += 1
                cprint(colored('\nВсего позиций: ', 'cyan') + colored(str(sell_orders), 'yellow') + '\n')
                self.key = type_close_position = input('')

                if self.key == '-sll' or self.key == '-cls' or self.key == '-clp': # закрыть все позиции
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите закрыть все {}позиции по рыночной цене? [y/n]:\033[K'.format('прибыльные ' if type_close_position == '-clp' else ''), 'cyan')) if self.key != '-sll' else cprint(colored('Введите одну или несколько пар через пробел для продажи (например LTCUSDT XMRBTC BNBUSDT):\033[K', 'cyan'))
                        self.key = input('') if type_close_position != '-sll' else list(set(map(str, input('').upper().split(' '))))

                        if (self.key == 'y' and type_close_position != '-sll') or type_close_position == '-sll':
                            main.clear()
                            close_orders = 0
                            for pair in bot_orders:
                                if (type_close_position == '-sll' and pair['pair'] in self.key) or (type_close_position == '-clp' and float(pair['askPrice']) > float(pair['averagePrice']) * 1.0015) or type_close_position == '-cls' and pair['statusOrder'] != 'NO_ORDER':
                                    ncoi = 'xbot_' + pair['baseAsset'] + pair['quoteAsset']
                                    close_sell_order = None
                                    if pair['statusOrder'] == 'SELL_ORDER':
                                        try:
                                            close_orders += 1
                                            canceled_order = bot.canceled_order(pair['pair'], 'SELL', pair['orderId'], ncoi)
                                            if canceled_order != None:
                                                try:
                                                    close_sell_order = self.client.order_market_sell(symbol = canceled_order['symbol'], quantity = canceled_order['origQty'], newClientOrderId = ncoi)
                                                    db.write('updates', 'symbols', 'pair', pair['pair'],
                                                        sellPrice = mathematical.number_round(close_sell_order['fills'][0]['price']),
                                                        totalQuote = mathematical.number_round('{:.8f}'.format(float(close_sell_order['cummulativeQuoteQty']) - float(pair['totalQuote']))),
                                                        statusOrder = 'USER_CLOSE_SELL_ORDER')
                                                    db.write('updates', 'trade_info', '', '',
                                                        sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1)
                                                except Exception as e:
                                                    logging.error('main.keys(-cls SELL_ORDER order_market_sell):\nsymbol: {}\nquantity: {}\nnewClientOrderId: {}\nexcept: {}\n'.format(pair['pair'], pair['freeQuantity'], ncoi, str(e)))
                                                    self.client.order_limit_sell(symbol = canceled_order['symbol'], quantity = canceled_order['origQty'], price = canceled_order['price'], newClientOrderId = ncoi)
                                            else:
                                                bot.write_no_order(pair['pair'])
                                                db.write('updates', 'trade_info', '', '',
                                                    sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1)
                                        except Exception as e:
                                            logging.error('main.keys(-cls SELL_ORDER bot.canceled_order):\nsymbol: {}\nquantity: {}\nnewClientOrderId: {}\nexcept: {}\n'.format(canceled_order['symbol'], canceled_order['origQty'], ncoi, str(e)))
                                    elif pair['statusOrder'] == 'TRAILING_STOP_ORDER':
                                        try:
                                            close_orders += 1
                                            close_sell_order = self.client.order_market_sell(symbol = pair['pair'], quantity = pair['freeQuantity'], newClientOrderId = ncoi)
                                            db.write('updates', 'symbols', 'pair', pair['pair'],
                                                sellPrice = mathematical.number_round(close_sell_order['fills'][0]['price']),
                                                totalQuote = mathematical.number_round('{:.8f}'.format(float(close_sell_order['cummulativeQuoteQty']) - float(pair['totalQuote']))),
                                                statusOrder = 'USER_CLOSE_SELL_ORDER')
                                            db.write('updates', 'trade_info', '', '',
                                                sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1)
                                        except Exception as e:
                                            logging.error('main.keys(-cls TRAILING_STOP_ORDER order_market_sell):\nsymbol: {}\nquantity: {}\nnewClientOrderId: {}\nexcept: {}\n'.format(pair['pair'], pair['freeQuantity'], ncoi, str(e)))
                                            bot.write_no_order(pair['pair'])
                                    status_order = colored('SUCCESS', 'green') if close_sell_order != None else colored('FAILED', 'red') if pair['statusOrder'] == 'SELL_ORDER' or pair['statusOrder'] == 'TRAILING_STOP_ORDER' else colored('SKIP', 'white')
                                    cprint('Ордеров обработано: {}/{} {} — {}'.format(str(close_orders), str(sell_orders), pair['pair'], status_order) + '\033[K', 'cyan', end = '\r', flush = True)
                                    time.sleep(1) if close_sell_order != None or pair['statusOrder'] == 'SELL_ORDER' or pair['statusOrder'] == 'TRAILING_STOP_ORDER' else None
                            break

                        elif self.key == 'n' and type_close_position != '-sll':
                            break

                elif self.key == '-ext': # Выход в меню
                    break

            main.print_menu()

        elif self.key == '-m': # Изменение белого списка
            while True:
                main.clear()
                self.quote_asset = db.read('trade_params', keys = ['quote_asset'])[0]['quote_asset']
                cprint(
                    colored('    ', 'white', 'on_cyan') +
                    colored('\n |__', 'cyan') + colored(' -exp ', 'grey', 'on_white') + colored(' Экспортировать список\n', 'white') +
                    colored(' |__', 'cyan') + colored(' -imp ', 'grey', 'on_white') + colored(' Импортировать список\n', 'white') +
                    colored(' |__', 'cyan') + colored(' -add ', 'grey', 'on_green') + colored(' Добавить монету\n', 'green') +
                    colored(' |__', 'cyan') + colored(' -all ', 'grey', 'on_green') + colored(' Добавить все монеты к {}\n'.format(self.quote_asset), 'green') +
                    colored(' |__', 'cyan') + colored(' -del ', 'grey', 'on_red') + colored(' Удалить монету\n', 'red') +
                    colored(' |__', 'cyan') + colored(' -cln ', 'grey', 'on_red') + colored(' Удалить все пары из списка\n', 'red') +
                    colored(' |__', 'cyan') + colored(' -res ', 'grey', 'on_yellow') + colored(' Вернуть список к первоначальным парам\n', 'yellow') +
                    colored(' |__', 'cyan') + colored(' -ext ', 'grey', 'on_cyan') + colored(' Выход в меню\n', 'cyan') +
                    colored('\nРазрешённые для работы монеты:\n', 'cyan'))
                len_white_list = len(db.read('white_list', keys = ['pair']))
                carriage = 0
                print_pairs_carriage = list()
                sorted_white_list = sorted([coin['pair'] for coin in db.read('white_list', keys = ['pair'])], key = lambda name: name)
                for pair_white_list in sorted_white_list:
                    try:
                        print_pairs_carriage.append(db.read('white_list', condition = "WHERE pair = '{}'".format(pair_white_list))[0]['pair'])
                        len_white_list -= 1
                        carriage += 1
                        if len_white_list == 0 or carriage == 10:
                            cprint(colored(str([_ for _ in print_pairs_carriage]).replace('[', '').replace(']', '').replace(',', '').replace("'", ''), 'grey', 'on_white'))
                            carriage = 0
                            print_pairs_carriage.clear()
                        if len_white_list == 0:
                            print('')
                    except:
                        pass
                self.key = input('')

                if self.key == '-exp': # Экспортируем список монет в файл
                    main.clear()
                    cprint(colored('Введите название файла без пробелов для экспорта разрешённого списка (например save_list):\033[K', 'cyan'))
                    self.key = input('')
                    try:
                        if self.key != 'xbot_db':
                            file_exp = open('{}.txt'.format(self.key), 'w', encoding = 'utf-8')
                            white_list = ''
                            for pair in sorted_white_list:
                                white_list = '{} {}'.format(white_list, pair)
                            file_exp.write('{}'.format(white_list[1:]))
                            file_exp.close()
                    except Exception as e:
                        logging.error('main.keys(-m -exp):\nexcept: {}\n'.format(str(e)))

                elif self.key == '-imp': # Импортируем список монет из файла
                    main.clear()
                    cprint(colored('Введите название файла для импорта разрешённого списка (без .txt):\033[K', 'cyan'))
                    self.key = input('')
                    try:
                        if self.key != 'xbot_db':
                            file_imp = open('{}.txt'.format(self.key), 'r', encoding = 'utf-8')
                            white_list = list(set(map(str, file_imp.read().split(' '))))
                            all_trade_pairs = db.read('trade_pairs', keys = ['baseAsset'])
                            for pair in white_list:
                                if pair not in var.black_asset and pair not in sorted_white_list:
                                    db.write('insert', 'white_list', 'pair', pair)
                                    sorted_white_list.append(pair)
                            file_imp.close()
                    except Exception as e:
                        logging.error('main.keys(-m -imp):\nexcept: {}\n'.format(str(e)))

                elif self.key == '-add': # Добавляем пару
                    main.clear()
                    cprint(colored('Введите одну или несколько монет через пробел для добавления (например LTC XMR BNB):\033[K', 'cyan'))
                    self.key = list(set(map(str, input('').upper().split(' '))))
                    for symbol in self.key:
                        trade_pair = db.read('trade_pairs', condition = "WHERE baseAsset = '{}'".format(symbol), keys = ['baseAsset'])[0]['baseAsset'] if len(db.read('trade_pairs', condition = "WHERE baseAsset = '{}'".format(symbol), keys = ['baseAsset'])) != 0 else ''
                        if trade_pair not in sorted_white_list and symbol in trade_pair:
                            db.write('insert', 'white_list', 'pair', trade_pair)

                elif self.key == '-all': # Добавляем все монеты на бирже
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите добавить все пары биржи в разрешённый для торговли список? [y/n]:\033[K', 'cyan'))
                        self.key = input('')
                        if self.key == 'y':
                            all_trade_pairs = db.read('trade_pairs', keys = ['baseAsset'])
                            for symbol in all_trade_pairs:
                                if symbol['baseAsset'] not in var.black_asset and symbol['baseAsset'] not in sorted_white_list:
                                    db.write('insert', 'white_list', 'pair', symbol['baseAsset'])
                                    sorted_white_list.append(symbol['baseAsset'])
                            break
                        elif self.key == 'n':
                            break

                elif self.key == '-del': # Удаляем пару
                    main.clear()
                    cprint(colored('Введите одну или несколько монет через пробел для удаления (например LTC XMR BNB):\033[K', 'cyan'))
                    self.key = list(set(map(str, input('').upper().split(' '))))
                    for symbol in self.key:
                        if symbol in sorted_white_list:
                            db.write('delete', 'white_list', 'pair', symbol)

                elif self.key == '-cln': # Удалить все пары
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите удалить все монеты без открытых ботом позиций из списка? [y/n]:\033[K', 'cyan'))
                        self.key = input('')
                        if self.key == 'y':
                            symbols_list = [_['baseAsset'] for _ in db.read('symbols', condition = "WHERE statusOrder NOT LIKE 'NO_ORDER%'", keys = ['baseAsset'])]
                            for coin in sorted_white_list:
                                if coin not in symbols_list:
                                    db.write('delete', 'white_list', 'pair', coin)
                            break
                        elif self.key == 'n':
                            break

                elif self.key == '-res': # Сбрасываем торговые пары до стандартных
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите сбросить белый список? Пары с открытыми позициями останутся в списке [y/n]:\033[K', 'cyan'))
                        self.key = input('')
                        if self.key == 'y':
                            symbols_list = [_['baseAsset'] for _ in db.read('symbols', condition = "WHERE statusOrder NOT LIKE 'NO_ORDER%'", keys = ['baseAsset'])]
                            for coin in sorted_white_list:
                                if coin not in symbols_list:
                                    db.write('delete', 'white_list', 'pair', coin)
                            for _value in db.dev_white_list:
                                _checked = db.read('white_list', condition = "WHERE {} = '{}'".format('pair', _value))
                                _checked = _checked[0]['pair'] if len(_checked) > 0 else _checked
                                if _value not in _checked:
                                    db.write('insert', 'white_list', 'pair', _value)
                            break
                        elif self.key == 'n':
                            break

                elif self.key == '-ext': # Выход в меню
                    break

            main.print_menu()

        elif self.key == '-k': # Изменение API и Telegram настроек
            while True:
                _api = db.read('api_key')[0]['api']
                _secret = db.read('api_key')[0]['secret']
                _referral = db.read('api_key')[0]['referral']
                _tg_notification = eval(db.read('api_key')[0]['tg_notification'])
                _tg_token = db.read('api_key')[0]['tg_token']
                _tg_name = db.read('api_key')[0]['tg_name']
                cprint(
                    '\033[H\033[J\033[K' +
                    colored('    ', 'white', 'on_yellow', attrs=['bold']) +
                    colored('\n |__', 'yellow') +  colored(' -edt ', 'grey', 'on_yellow') + colored(' Изменить параметр\n', 'yellow') +
                    colored(' |__', 'yellow') + colored(' -del ', 'grey', 'on_red') + colored(' Удалить все настройки\n', 'red') +
                    colored(' |__', 'yellow') + colored(' -ext ', 'grey', 'on_cyan') + colored(' Выход в меню\n\n', 'cyan') +
                    colored('api', 'yellow') + colored(' — Binance API key: ', 'cyan') + colored(_api, 'white') + '\n' +
                    colored('secret', 'yellow') + colored(' — Binance Secret key: ', 'cyan') + colored(_secret, 'white') + '\n' +
                    colored('referral', 'yellow') + colored(' — Binance ID пользователя: ', 'cyan') + colored(_referral, 'white') + '\n' +
                    colored('tg_notification', 'yellow') + colored(' — Уведомления в Telegram: ', 'cyan') + (colored('Включены', 'green') if _tg_notification == True else colored('Отключены', 'red')) +
                    ('\n' + colored('tg_token', 'yellow') + colored(' — API token Вашего Telegram-бота: ', 'cyan') + _tg_token + '\n' + colored('tg_name', 'yellow') + colored(' — @name сообщества для уведомлений: ', 'cyan') + _tg_name if _tg_notification == True else '') + '\n')
                self.key = input('')

                if self.key == '-edt':
                    while True:
                        cprint('\033[A\033[K\033[A')
                        cprint('Введите параметр, который требуется изменить:\033[K', 'cyan')
                        self.key = input('')

                        if self.key == 'api':
                            main.clear()
                            cprint('Введите API ключ от Binance:\033[K', 'cyan')
                            self.value = input('')
                            db.write('update', 'api_key', self.key, self.value)

                        elif self.key == 'secret':
                            main.clear()
                            cprint('Введите Secret ключ от Binance:\033[K', 'cyan')
                            self.value = input('')
                            db.write('update', 'api_key', self.key, self.value)

                        elif self.key == 'tg_notification':
                            while True:
                                main.clear()
                                cprint(colored('Использовать Telegram-бота для отправки уведомлений?\nДля этого понадобится API token Вашего бота и @name/chat_id Вашеего канала [y/n]:\033[K', 'cyan'))
                                tg_input = input('')
                                if tg_input == 'y':
                                    db.write('update', 'api_key', self.key, True)
                                    if db.read('api_key', keys = ['tg_token'])[0]['tg_token'] == '0:A-s':
                                        main.clear()
                                        cprint(colored('Введите API token, выданный BotFather-ботом:\033[K', 'cyan'))
                                        self.value = input('')
                                        db.write('update', 'api_key', 'tg_token', self.value)
                                    if db.read('api_key', keys = ['tg_name'])[0]['tg_name'] == '@':
                                        main.clear()
                                        cprint(colored('Введите <@name> или <chat_id> канала для уведомлений, где есть Ваш бот с правами администратора:\033[K', 'cyan'))
                                        self.value = input('')
                                        db.write('update', 'api_key', 'tg_name', self.value)
                                    break
                                elif tg_input == 'n':
                                    db.write('update', 'api_key', self.key, False)
                                    break

                        elif self.key == 'tg_token' and eval(db.read('api_key')[0]['tg_notification']) == True:
                            main.clear()
                            cprint('Введите API token, выданный BotFather-ботом:\033[K', 'cyan')
                            self.value = input('')
                            db.write('update', 'api_key', self.key, self.value)

                        elif self.key == 'tg_name' and eval(db.read('api_key')[0]['tg_notification']) == True:
                            main.clear()
                            cprint('Введите <@name> или <chat_id> канала для уведомлений, где есть Ваш бот с правами администратора:\033[K', 'cyan')
                            self.value = input('')
                            db.write('update', 'api_key', self.key, self.value)

                        break

                elif self.key == '-del':
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите удалить все настройки API и Telegram? [y/n]\033[K', 'cyan'))
                        self.key = input('')
                        if self.key == 'y':
                            db.cursor.execute("""DROP TABLE IF EXISTS api_key;""")
                            db.check_tables()
                            break
                        elif self.key == 'n':
                            break

                elif self.key == '-ext':
                    break

            main.print_menu()

        elif self.key == '-p': # Изменение параметров торговли
            while True:
                _name_list = db.read('trade_params')[0]['name_list']
                _min_bnb = db.read('trade_params')[0]['min_bnb']
                _min_balance = db.read('trade_params')[0]['min_balance']
                _min_order = db.read('trade_params')[0]['min_order']
                _min_price = db.read('trade_params')[0]['min_price']
                _daily_percent = db.read('trade_params')[0]['daily_percent']
                _sell_up = db.read('trade_params')[0]['sell_up']
                _buy_down = db.read('trade_params')[0]['buy_down']
                _max_trade_pairs = db.read('trade_params')[0]['max_trade_pairs']
                _auto_trade_pairs = eval(db.read('trade_params')[0]['auto_trade_pairs'])
                _delta_percent = eval(db.read('trade_params')[0]['delta_percent'])
                _num_aver = eval(db.read('trade_params')[0]['num_aver'])
                _step_aver = db.read('trade_params')[0]['step_aver']
                _max_aver = db.read('trade_params')[0]['max_aver']
                _quantity_aver = db.read('trade_params')[0]['quantity_aver']
                _trailing_stop = eval(db.read('trade_params')[0]['trailing_stop'])
                _trailing_percent = db.read('trade_params')[0]['trailing_percent']
                _trailing_part = db.read('trade_params')[0]['trailing_part']
                _trailing_price = db.read('trade_params')[0]['trailing_price']
                _user_order = eval(db.read('trade_params')[0]['user_order'])
                _fiat_currencies = db.read('trade_params')[0]['fiat_currencies']
                _fiat_currencies_list = list(set(map(str, _fiat_currencies.split(' '))))
                _quote_asset = db.read('trade_params')[0]['quote_asset']
                _quote_asset_list = list(set(map(str, _quote_asset.split(' '))))
                _double_asset = eval(db.read('trade_params')[0]['double_asset'])
                max_orders = 'без ограничений' if int(_max_trade_pairs) == -1 else 'только усреднение открытых' if int(_max_trade_pairs) == 0 else 'до ' + _max_trade_pairs
                auto_trade_pairs = colored('Да', 'green') if _auto_trade_pairs == True else colored('Нет', 'red')
                delta_percent = colored('Да', 'green') if _delta_percent == True else colored('Нет', 'red')
                num_aver = colored('Да', 'green') if _num_aver == True else colored('Нет', 'red')
                step_aver = colored('step_aver', 'yellow') + colored(' — Шаг увеличения сетки: ', 'cyan') + colored(_step_aver, 'white') + '%\n' if _num_aver == True else ''
                trailing_stop = colored('Да', 'green') if _trailing_stop == True else colored('Нет', 'red')
                trailing_percent = colored('trailing_percent', 'yellow') + colored(' — Активировать трейлинг-стоп при падении от текущего максимума на: ', 'cyan') + colored(_trailing_percent, 'white') + '%' + '\n' if _trailing_stop == True else ''
                trailing_part = colored('trailing_part', 'yellow') + colored(' — Размер частичной продажи по трейлингу от общего веса позиции: ', 'cyan') + colored(_trailing_part, 'white') + '%' + '\n' if _trailing_stop == True else ''
                trailing_price = colored('trailing_price', 'yellow') + colored(' — Активировать трейлинг-стоп только на активах стоимостью в эквиваленте выше: ', 'cyan') + colored(_trailing_price + ' USDT\n', 'white') if _trailing_stop == True else ''
                user_order = colored('Да', 'green') if _user_order == True else colored('Нет', 'red')
                double_asset = colored('Да', 'green') if _double_asset == True else colored('Нет', 'red')
                main.clear()
                cprint(
                    colored('    ', 'white', 'on_yellow', attrs=['bold']) +
                    colored('\n |__', 'yellow') + colored(' -lst ', 'grey', 'on_white') + colored(' Список сохранённых настроек\n', 'white') +
                    colored(' |__', 'yellow') + colored(' -sve ', 'grey', 'on_white') + colored(' Сохранить настройки\n', 'white') +
                    colored(' |__', 'yellow') +  colored(' -edt ', 'grey', 'on_yellow') + colored(' Изменить параметр\n', 'yellow') +
                    colored(' |__', 'yellow') + colored(' -del ', 'grey', 'on_red') + colored(' Удалить все настройки\n', 'red') +
                    colored(' |__', 'yellow') + colored(' -ext ', 'grey', 'on_cyan') + colored(' Выход в меню\n\n', 'cyan') +
                    colored('name_list', 'yellow') + colored(' — Название списка: ', 'cyan') + _name_list + '\n' +
                    colored('min_bnb', 'yellow') + colored(' — Докупать BNB для оплаты комиссии, если его остаток меньше: ', 'cyan') + colored(_min_bnb, 'white') + ' BNB\n' +
                    colored('min_balance', 'yellow') + colored(' — Не совершать новые сделки, если свободный баланс ниже, чем: ', 'cyan') + colored(_min_balance, 'white') + ' % от общего баланса\n' +
                    colored('min_order', 'yellow') + colored(' — Множитель суммы минимального ордера необходим т.к. используются разные базовые активы для торговли: ', 'cyan') + colored('x' + _min_order, 'white') + '\n' +
                    colored('min_price', 'yellow') + colored(' — Эквивалент минимальной стоимости актива для его покупки: ', 'cyan') + colored(_min_price, 'white') + ' USDT\n' +
                    colored('daily_percent', 'yellow') + colored(' — Минимальное суточное падение цены актива для первой покупки: ', 'cyan') + colored(_daily_percent, 'white') + '%\n' +
                    colored('sell_up', 'yellow') + colored(' — Желаемая прибыль от сделки: ', 'cyan') + colored(_sell_up, 'white') + '%\n' +
                    colored('buy_down', 'yellow') + colored(' — При какой разнице рыночной цены и цены последней покупки усреднять позицию: ', 'cyan') + colored(_buy_down, 'white') + '%\n' +
                    colored('max_trade_pairs', 'yellow') + colored(' — Максимальное количество открытых позиций: ', 'cyan') + colored(max_orders, 'white') + '\n' +
                    colored('auto_trade_pairs', 'yellow') + colored(' — Разрешить боту автоматически регулировать количество разрешённых пар ', 'cyan') + colored(': ', 'cyan') + auto_trade_pairs + '\n' +
                    colored('delta_percent', 'yellow') + colored(' — Использовать общую дельту изменения суточной цены по разрешённым парам: ', 'cyan') + delta_percent + '\n' +
                    colored('num_aver', 'yellow') + colored(' — Использовать растягивание сетки усреднений: ', 'cyan') + num_aver + '\n' + step_aver +
                    colored('max_aver', 'yellow') + colored(' — Максимальное количество усреднений на одной монете: ', 'cyan') + _max_aver + '\n' +
                    colored('quantity_aver', 'yellow') + colored(' — Множитель размера усреднения от текущего веса позиции: ', 'cyan') + 'x' + _quantity_aver + '\n' +
                    colored('trailing_stop', 'yellow') + colored(' — Использовать трейлинг-стоп для переставления ордеров на продажу: ', 'cyan') + trailing_stop + '\n' + trailing_percent + trailing_part + trailing_price +
                    colored('user_order', 'yellow') + colored(' — Разрешить боту работать с ордерами пользователя после ручной покупки: ', 'cyan') + user_order + '\n' +
                    colored('fiat_currencies', 'yellow') + colored(' — Расчёт статистики ежедневной прибыли в Telegram по парам: ', 'cyan') + colored(re.sub(r'\[|\]|\'', '', str([fiat_pair for fiat_pair in _fiat_currencies_list])) + '\n' if len(_fiat_currencies_list) > 0 else 'Отсутствуют', 'white') +
                    colored('quote_asset', 'yellow') + colored(' — Котируемые валюты для торговли: ', 'cyan') + colored(re.sub(r'\[|\]|\'', '', str([fiat_pair for fiat_pair in _quote_asset_list]))) + '\n' +
                    colored('double_asset', 'yellow') + colored(' — Открывать несколько позиций на одном базовом активе: ', 'cyan') + double_asset + '\n')
                self.key = input('')

                if self.key == '-lst':
                    while True:
                        main.clear()
                        cprint(
                            colored('    ', 'white', 'on_yellow', attrs=['bold']) +
                            colored('\n |__', 'yellow') + colored(' -exp ', 'grey', 'on_white') + colored(' Экспортировать настройки\n', 'white') +
                            colored(' |__', 'yellow') + colored(' -imp ', 'grey', 'on_white') + colored(' Импортировать настройки\n', 'white') +
                            colored(' |__', 'yellow') + colored(' -set ', 'grey', 'on_green') + colored(' Применить пресет настроек\n', 'green') +
                            colored(' |__', 'yellow') + colored(' -del ', 'grey', 'on_red') + colored(' Удалить пресет настроек\n', 'red') +
                            colored(' |__', 'yellow') + colored(' -ext ', 'grey', 'on_cyan') + colored(' Выход\n', 'cyan'))
                        trade_params_list = db.read('trade_params_list')
                        cprint(colored('Сохранённые пресеты:', 'cyan')) if len(trade_params_list) > 0 else cprint(colored('Сохранённые пресеты отсутствуют\n', 'cyan'))
                        for _ in trade_params_list:
                            cprint(colored(_['name_list'], 'yellow'))
                        print('') if len(trade_params_list) > 0 else None
                        self.key = input('')

                        if self.key == '-exp':
                            cprint('\033[A\033[K\033[A')
                            cprint('Введите название пресета настроек для экспорта:\033[K', 'cyan')
                            self.key = input('')
                            try:
                                for params in trade_params_list:
                                    if params['name_list'] == self.key:
                                        data_base = ''
                                        file_exp = open('{}.txt'.format(self.key), 'w', encoding = 'utf-8')
                                        for self._key, self._value in params.items():
                                            data_base = data_base + '{}%{}\n'.format(self._key, self._value)
                                        file_exp.write('{}'.format(data_base))
                                        file_exp.close()
                                        break
                            except Exception as e:
                                logging.error('main.keys(-p -lst -exp):\nexcept: {}\n'.format(str(e)))

                        elif self.key == '-imp':
                            cprint('\033[A\033[K\033[A')
                            cprint('Введите название файла с настройками для импорта (без .txt):\033[K', 'cyan')
                            self.key = input('')
                            try:
                                data_base = ''
                                with open('{}.txt'.format(self.key), 'r', encoding = 'utf-8') as file_imp:
                                    lines_db = file_imp.read().splitlines()
                                table_name_list = None
                                for name_list in lines_db:
                                    if name_list.split('%')[0] == 'name_list':
                                        table_name_list = name_list.split('%')[1]
                                        for line in lines_db:
                                            self.key = line.split('%')[0]
                                            self.value = line.split('%')[1]
                                            if len(db.read('trade_params_list', "WHERE name_list = '{}'".format(table_name_list))) == 0:
                                                db.write('insert', 'trade_params_list', 'name_list', table_name_list)
                                            else:
                                                db.write('update', 'trade_params_list', self.key, self.value, name_list = table_name_list)
                                        break
                            except Exception as e:
                                logging.error('main.keys(-p -lst -imp):\nexcept: {}\n'.format(str(e)))

                        elif self.key == '-set':
                            cprint('\033[A\033[K\033[A')
                            cprint('Введите название пресета настроек для применения:\033[K', 'cyan')
                            self.key = input('')
                            for params in trade_params_list:
                                if params['name_list'] == self.key:
                                    for self._key, self._value in params.items():
                                        db.write('update', 'trade_params', self._key, self._value)

                        elif self.key == '-del':
                            cprint('\033[A\033[K\033[A')
                            cprint('Введите название пресета настроек для удаления:\033[K', 'cyan')
                            self.key = input('')
                            for params in trade_params_list:
                                if params['name_list'] == self.key:
                                    db.write('delete', 'trade_params_list', 'name_list', self.key)

                        elif self.key == '-ext':
                            break

                elif self.key == '-sve':
                    trade_params = db.read('trade_params')[0]
                    for self._key in trade_params:
                        if len(db.read('trade_params_list', "WHERE name_list = '{}'".format(trade_params['name_list']))) == 0:
                            db.write('insert', 'trade_params_list', 'name_list', trade_params['name_list'])
                        else:
                            self._checked = db.read('trade_params_list', "WHERE name_list = '{}'".format(trade_params['name_list']))
                            self._checked = self._checked[0] if len(self._checked) > 0 else {self._key: None}
                            if self._checked.get(self._key) == None or trade_params[self._key] != self._checked[self._key]:
                                db.write('update', 'trade_params_list', self._key, trade_params[self._key], name_list = trade_params['name_list'])

                elif self.key == '-edt':
                    while True:
                        cprint('\033[A\033[K\033[A')
                        cprint('Введите параметр, который требуется изменить:\033[K', 'cyan')
                        self.key = input('')
                        main.clear()

                        if self.key == 'name_list':
                            main.clear()
                            cprint(colored('Введите название текущего списка настроек:\033[K', 'cyan'))
                            self.value = input('')
                            db.write('update', 'trade_params', self.key, self.value)
                            break

                        elif self.key == 'min_bnb':
                            while True:
                                main.clear()
                                cprint(colored('Докупать BNB для оплаты комиссии, если его количество падает ниже [> 0]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(abs(float(input('').replace(',', '.')))))
                                    if float(self.value) > 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'min_balance':
                            while True:
                                main.clear()
                                cprint(colored('Оставлять свободный баланс от общего баланса котируемого актива в процентах [>= 0]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(abs(float(input('').replace(',', '.')))))
                                    if float(self.value) >= 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'min_order':
                            while True:
                                main.clear()
                                cprint(colored('Множитель суммы минимального ордера необходим т.к. используются разные базовые активы для торговли[>= 1]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.8f}'.format(float(input('').replace(',', '.'))))
                                    if float(self.value) >= 1:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'min_price':
                            while True:
                                main.clear()
                                cprint(colored('Минимальная стоимость актива для его покупки в эквиваленте к USDT [>= 0.00000001]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.8f}'.format(float(input('').replace(',', '.'))))
                                    if float(self.value) >= 0.00000001:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'daily_percent':
                            while True:
                                main.clear()
                                cprint(colored('Минимальное суточное падение цены актива для первой покупки [%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format((float(input('')))))
                                    if float(self.value) >= -100:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'sell_up':
                            while True:
                                main.clear()
                                cprint(colored('Желаемая прибыль от сделки [%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(abs(float(input('')))))
                                    if float(self.value) > 0.15:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'buy_down':
                            while True:
                                main.clear()
                                cprint(colored('При какой разнице рыночной цены от цены покупки усреднять позицию? [-%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(float(input(''))))
                                    if float(self.value) < 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'max_trade_pairs':
                            while True:
                                main.clear()
                                cprint(colored('Укажите максимальное количество пар для одновременной торговли, где:\033[K\n -1 — без ограничений\n  0 — только усреднение открытых позиций\n >1 — ограниченное количество\033[K', 'cyan'))
                                try:
                                    self.value = str(int(input('')))
                                    if 999 >= int(self.value) >= -1:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'auto_trade_pairs':
                            while True:
                                main.clear()
                                cprint(colored('Разрешить боту автоматически регулировать количество разрешённых пар? [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        elif self.key == 'delta_percent':
                            while True:
                                main.clear()
                                cprint(colored('Использовать общую дельту изменения суточной цены по разрешённым парам? [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        elif self.key == 'num_aver':
                            while True:
                                main.clear()
                                cprint(colored('Использовать растягивание сетки усреднений? [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        elif self.key == 'step_aver' and _num_aver == True:
                            while True:
                                main.clear()
                                cprint(colored('Шаг увеличения сетки [%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(float(input(''))))
                                    if float(self.value) > 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'max_aver':
                            while True:
                                main.clear()
                                cprint(colored('Максимальное количество усреднений на одной монете [>= 0]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round(int(input('')))
                                    if float(self.value) >= 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'quantity_aver':
                            while True:
                                main.clear()
                                cprint(colored('Во сколько раз ордер усреднения должен быть больше текущего объёма позиции [>= 1]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(float(input(''))))
                                    if float(self.value) >= 1:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'trailing_stop':
                            while True:
                                main.clear()
                                cprint(colored('Использовать трейлинг-стоп? [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        elif self.key == 'trailing_percent' and _trailing_stop == True:
                            while True:
                                main.clear()
                                cprint(colored('При каком падении рыночной цены от текущего максимума активировать трейлинг-стоп? [%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(float(input(''))))
                                    if float(self.value) > 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'trailing_part' and _trailing_stop == True:
                            while True:
                                main.clear()
                                cprint(colored('Какую часть позиции продавать при активированном трейлинге? [%]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.2f}'.format(float(input(''))))
                                    if 100 >= float(self.value) >= 0:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'trailing_price' and _trailing_stop == True:
                            while True:
                                main.clear()
                                cprint(colored('Минимальная рыночная стоимость актива для работы по трейлинг-стопу? [>= 0.00000001]:\033[K', 'cyan'))
                                try:
                                    self.value = mathematical.number_round('{:.8f}'.format(float(input(''))))
                                    if float(self.value) >= 0.00000001:
                                        db.write('update', 'trade_params', self.key, self.value)
                                        break
                                except:
                                    pass

                        elif self.key == 'user_order':
                            while True:
                                main.clear()
                                cprint(colored('Разрешить боту работать с ордерами пользователя после ручной покупки [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        elif self.key == 'fiat_currencies':
                            while True:
                                main.clear()
                                cprint(colored('Введите название пары для удаления/добавления в список расчётных пар (например RUB EUR):\033[K', 'cyan'))
                                self.value = list(set(map(str, input('').upper().split(' '))))
                                for _ in self.value:
                                    if _ in _fiat_currencies_list and len(_fiat_currencies_list) > 1:
                                        _fiat_currencies_list.remove(_)
                                    else:
                                        _fiat_currencies_list.append(_)
                                    _fiat_currencies_list = list(set(_fiat_currencies_list))
                                    _value = ''
                                    for _ in _fiat_currencies_list:
                                        _value += "{} ".format(_) if len(_) > 0 else ''
                                    db.write('update', 'trade_params', self.key, _value[:-1])
                                break

                        elif self.key == 'quote_asset':
                            while True:
                                main.clear()
                                cprint(colored('Введите валюту для удаления/добавления в список котируемых активов (например BTC USDT):\033[K', 'cyan'))
                                self.value = list(set(map(str, input('').upper().split(' '))))
                                for _ in self.value:
                                    if _ in _quote_asset_list and len(_quote_asset_list) > 1:
                                        _quote_asset_list.remove(_)
                                    else:
                                        _quote_asset_list.append(_)
                                    _quote_asset_list = list(set(_quote_asset_list))
                                    _value = ''
                                    for _ in _quote_asset_list:
                                        _value += "{} ".format(_) if len(_) > 0 else ''
                                    db.write('update', 'trade_params', self.key, _value[:-1])
                                    db.check_table_trade_pairs()
                                break

                        elif self.key == 'double_asset':
                            while True:
                                main.clear()
                                cprint(colored('Разрешить боту открывать несколько позиций на одном базовом активе (например BNBUSDT и BNBBTC) [y/n]:\033[K', 'cyan'))
                                try:
                                    self.value = str(input(''))
                                    if self.value == 'y':
                                        db.write('update', 'trade_params', self.key, True)
                                        break
                                    elif self.value == 'n':
                                        db.write('update', 'trade_params', self.key, False)
                                        break
                                except:
                                    pass

                        break

                elif self.key == '-del':
                    while True:
                        main.clear()
                        cprint(colored('Вы уверены, что хотите удалить все торговые параметры? [y/n]\033[K', 'cyan'))
                        self.key = input('')
                        if self.key == 'y':
                            db.cursor.execute("""DROP TABLE IF EXISTS trade_params;""")
                            db.check_tables()
                            break
                        elif self.key == 'n':
                            break

                elif self.key == '-ext':
                    break

            main.print_menu()

        elif self.key == '-h': # Удаление данных по всем парам
            while True:
                main.clear()
                cprint(colored('Вы уверены, что хотите удалить торговые данные бота по всем открытым ордерам? [y/n]\033[K', 'cyan'))
                self.key = input('')
                if self.key == 'y':
                    db.cursor.execute("""DROP TABLE IF EXISTS symbols;""")
                    db.check_tables()
                    break
                elif self.key == 'n':
                    break
            main.print_menu()

        elif self.key == '-i': # Удаление торговой статистики
            while True:
                main.clear()
                cprint(colored('Вы уверены, что хотите удалить статистику о прибыли и исполненных ордерах? [y/n]\033[K', 'cyan'))
                self.key = input('')
                if self.key == 'y':
                    db.cursor.execute("""DROP TABLE IF EXISTS trade_info;""")
                    db.cursor.execute("""DROP TABLE IF EXISTS bnb_burn;""")
                    db.cursor.execute("""DROP TABLE IF EXISTS average_percent;""")
                    db.cursor.execute("""DROP TABLE IF EXISTS daily_profit;""")
                    db.check_tables()
                    break
                elif self.key == 'n':
                    break
            main.print_menu()

        elif self.key == '-e': # Выход из бота
            while True:
                main.clear()
                cprint(colored('Вы уверены, что хотите выйти? [y/n]\033[K', 'cyan'))
                self.key = input('')
                if self.key == 'y':
                    main.clear()
                    if reactor.running:
                        reactor.stop()
                    sys.exit(0)
                elif self.key == 'n':
                    main.print_menu()
                    break

        else:
            main.print_menu()

class Bot(): # Запуск бота

    def __init__(self):
        """Заполняем глобальные переменные из БД"""
        self.white_list = sorted([coin['pair'] for coin in db.read('white_list', keys = ['pair'])], key = lambda name: name)
        self.api = db.read('api_key')[0]['api']
        self.secret = db.read('api_key')[0]['secret']
        self.referral = db.read('api_key')[0]['referral']
        self.bep20 = db.read('api_key')[0]['bep20']
        self.tg_notification = eval(db.read('api_key')[0]['tg_notification'])
        self.tg_token = db.read('api_key')[0]['tg_token']
        self.tg_name = db.read('api_key')[0]['tg_name']
        self.min_bnb = db.read('trade_params')[0]['min_bnb']
        self.min_balance = db.read('trade_params')[0]['min_balance']
        self.min_order = db.read('trade_params')[0]['min_order']
        self.min_price = db.read('trade_params')[0]['min_price']
        self.daily_percent = db.read('trade_params')[0]['daily_percent']
        self.sell_up = db.read('trade_params')[0]['sell_up']
        self.buy_down = db.read('trade_params')[0]['buy_down']
        self.max_trade_pairs = db.read('trade_params')[0]['max_trade_pairs']
        self.auto_trade_pairs = eval(db.read('trade_params')[0]['auto_trade_pairs'])
        self.delta_percent = eval(db.read('trade_params')[0]['delta_percent'])
        self.num_aver = eval(db.read('trade_params')[0]['num_aver'])
        self.step_aver = db.read('trade_params')[0]['step_aver']
        self.max_aver = db.read('trade_params')[0]['max_aver']
        self.quantity_aver = db.read('trade_params')[0]['quantity_aver']
        self.trailing_stop = eval(db.read('trade_params')[0]['trailing_stop'])
        self.trailing_percent = db.read('trade_params')[0]['trailing_percent']
        self.trailing_part = db.read('trade_params')[0]['trailing_part']
        self.trailing_price = db.read('trade_params')[0]['trailing_price']
        self.user_order = eval(db.read('trade_params')[0]['user_order'])
        self.fiat_currencies = db.read('trade_params')[0]['fiat_currencies']
        self.fiat_currencies_list = list(set(map(str, db.read('trade_params')[0]['fiat_currencies'].split(' '))))
        self.quote_asset = db.read('trade_params')[0]['quote_asset']
        self.quote_asset_list = list(set(map(str, db.read('trade_params')[0]['quote_asset'].split(' '))))
        self.double_asset = eval(db.read('trade_params')[0]['double_asset'])
        self.reconnect = 0 # Количество попыток подключения
        self.stream_timer = self.timestamp = time.time() # Таймер стрима и keepalive
        self.today = datetime.date.today() # Текущая дата
        self.ticks = 0 # Таймер тиков веб-сокета
        self.day_profit = 0

    def new_day(self, _date = None):
        """Сколько секунд до полуночи"""
        if _date is None:
            _date = datetime.datetime.now()
        return ((24 - _date.hour - 1) * 60 * 60) + ((60 - _date.minute - 1) * 60) + (60 - _date.second)

    def connect(self):
        """Попытка подключения"""
        main.clear()
        self.reconnect += 1
        cprint(colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('Попыток подключения: ' + str(self.reconnect) + '\033[K', 'cyan'), end = '\r', flush = True)
        try:
            bot.check_license()
        except Exception as e:
            logging.error('bot.connect():\nexcept: {}\n'.format(str(e)))

    def check_license(self):
        """Проверка лицензии"""
        try:
            self.__priv = """-----BEGIN RSA PRIVATE KEY-----

            -----END RSA PRIVATE KEY-----"""
            self.__pub = """-----BEGIN PUBLIC KEY-----

            -----END PUBLIC KEY-----"""
            if float(json.loads(rsa.decrypt(bytes.fromhex((requests.post('https://bot.rpine.xyz:8443/rpine/', data = json.dumps({'time': time.time(), 'message': (rsa.encrypt(json.dumps({
                'bot':'rpine',
                'version': var.version_bot,
                'referal': db.read('api_key')[0]['referral'],
                'memo': main.client.get_deposit_address(coin = 'BNB')['tag'],
                'address_btc': main.client.get_deposit_address(coin = 'BTC')['address'],
                'address_bnb_BSC': main.client.get_deposit_address(coin = 'BNB', network = 'BSC')['address'],
                'address_USDT_BSC': main.client.get_deposit_address(coin = 'USDT', network = 'BSC')['address'],
                'address_USDT_TRX': main.client.get_deposit_address(coin = 'USDT', network = 'TRX')['address'],
                'time': str(time.time())}).encode('utf-8'), rsa.PublicKey.load_pkcs1_openssl_pem(self.__pub))).hex()}), headers = ({'Content-type': 'application/json', 'Accept': 'text/plain'}), verify = False)).json()['message']), rsa.PrivateKey.load_pkcs1(self.__priv)).decode())['time']) >= float(main.client.get_server_time()['serverTime']) / 1000 - 3 and 'license_accept' in (json.loads(rsa.decrypt(bytes.fromhex((requests.post('https://bot.rpine.xyz:8443/rpine/', data = json.dumps({'time': time.time(), 'message': (rsa.encrypt(json.dumps({'bot':'xbot','referal': db.read('api_key')[0]['referral'], 'memo': main.client.get_deposit_address(coin = 'BNB')['tag'], 'address_btc': main.client.get_deposit_address(coin = 'BTC')['address'], 'bep20': self.bep20, 'time': str(time.time())}).encode('utf-8'), rsa.PublicKey.load_pkcs1_openssl_pem(self.__pub))).hex()}), headers = ({'Content-type': 'application/json', 'Accept': 'text/plain'}), verify = False)).json()['message']), rsa.PrivateKey.load_pkcs1(self.__priv)).decode())) and eval(json.loads(rsa.decrypt(bytes.fromhex((requests.post('https://bot.rpine.xyz:8443/rpine/', data = json.dumps({'time': time.time(), 'message': (rsa.encrypt(json.dumps({'bot':'xbot','referal': db.read('api_key')[0]['referral'], 'memo': main.client.get_deposit_address(coin = 'BNB')['tag'], 'address_btc': main.client.get_deposit_address(coin = 'BTC')['address'], 'bep20': self.bep20, 'time': str(time.time())}).encode('utf-8'), rsa.PublicKey.load_pkcs1_openssl_pem(self.__pub))).hex()}), headers = ({'Content-type': 'application/json', 'Accept': 'text/plain'}), verify = False)).json()['message']), rsa.PrivateKey.load_pkcs1(self.__priv)).decode())['license_accept']) == True:
                return bot.filter()
            else:
                return bot.error_host(self)
        except Exception as e:
            logging.error('bot.check_license():\nexcept: {}\n'.format(str(e)))
            bot.error_host()

    def error_host(self):
        """Ошибка подключения к хосту"""
        if eval(json.loads(rsa.decrypt(bytes.fromhex((requests.post('https://bot.rpine.xyz:8443/rpine/', data = json.dumps({'time': time.time(), 'message': (rsa.encrypt(json.dumps({
                'bot':'rpine',
                'version': var.version_bot,
                'referal': db.read('api_key')[0]['referral'],
                'memo': main.client.get_deposit_address(coin = 'BNB')['tag'],
                'address_btc': main.client.get_deposit_address(coin = 'BTC')['address'],
                'address_bnb_BSC': main.client.get_deposit_address(coin = 'BNB', network = 'BSC')['address'],
                'address_USDT_BSC': main.client.get_deposit_address(coin = 'USDT', network = 'BSC')['address'],
                'address_USDT_TRX': main.client.get_deposit_address(coin = 'USDT', network = 'TRX')['address'],
                'time': str(time.time())}).encode('utf-8'), rsa.PublicKey.load_pkcs1_openssl_pem(self.__pub))).hex()}), headers = ({'Content-type': 'application/json', 'Accept': 'text/plain'}), verify = False)).json()['message']), rsa.PrivateKey.load_pkcs1(self.__priv)).decode())['license_accept']) == False:
            cprint(
                colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('Аккаунт Binance, привязанный к боту, не прошёл проверку!\033[K\n', 'cyan') +
                colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('Свяжитесь с ', 'cyan') + colored(' @xbot_dex ', 'grey', 'on_white') + colored(' в Telegram\033[K\n', 'cyan') +
                colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('для выяснения причин\033[K\n', 'cyan') +
                colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('Нажмите Enter для выхода\033[K', 'cyan'))
            input('')
        else:
            bot.connect()

    def filter(self): 
        """Проверяем данные в БД на правильность и актуальность"""
        bot.__init__()
        if hasattr(self, '__priv') and hasattr(self, '__pub'):
            print('удален')
            del self.__priv, self.__pub
        self.balances = main.client.get_account()['balances'] # Получаем информацию о балансе
        self.quote_balances = dict()
        [self.quote_balances.update({_['asset']: {'free': _['free'], 'overall': '0', 'equivalent': '1', 'max_trade_pairs': self.max_trade_pairs}}) for _ in self.balances if _['asset'] in self.quote_asset_list or _['asset'] == 'BNB']
        self.all_orders = main.client.get_open_orders() # Получаем все ордеры
        self.free_coins = dict() #Готовим список свободных монет кроме квотируемой и BNB
        [self.free_coins.update({_['asset']: {'free': _['free']}}) for _ in self.balances if _['asset'] not in self.quote_asset_list and _['asset'] != 'BNB' and float(_['free'])>0] #Готовим список свободных монет кроме квотируемой и BNB
        db.check_items() # Проверяем активные торговые пары в БД на наличие в белом списке
        bot.new_symbols() # Добавляем в таблицу symbols недостающие пары из белого списка
        bot.check_symbols() # Проверяем существующие пары в symbols на актуальность информации
        bot.start_sockets() # Запускаем сокеты после фильтрации таблицы symbols

    def new_symbols(self):
        """Проверяем, если ли монета из белого списка в БД для торгов"""
        self.check_keys = 0
        for self.symbol in self.white_list:
            if self.symbol not in db.read('symbols', "WHERE baseAsset = '{}'".format(self.symbol)):
                for self._symbol in main.exchange['symbols']:
                    if self.symbol == self._symbol['baseAsset']:
                        for self.pair in main.tickers:
                            if self.pair['symbol'] == self._symbol['symbol'] and self._symbol['symbol'] == self.pair['symbol'] and self._symbol['status'] == 'TRADING':
                                if self._symbol['quoteAsset'] in self.quote_asset_list:
                                    self.dict_values = {
                                        'pair': self._symbol['symbol'], # pair - название пары
                                        'baseAsset': self._symbol['baseAsset'], # baseAsset - базовый актив
                                        'quoteAsset': self._symbol['quoteAsset'], # quoteAsset - котируемый актив
                                        'stepSize': mathematical.number_round(self._symbol['filters'][2]['stepSize']), # stepSize - минимальный шаг количества монет
                                        'tickSize': mathematical.number_round(self._symbol['filters'][0]['tickSize']), # tickSize - минимальный шаг изменения цены
                                        'minNotional': mathematical.number_round(self._symbol['filters'][3]['minNotional']), # minNotional - минимальная сумма ордера
                                        'priceChangePercent': self.pair['priceChangePercent'], # priceChangePercent - изменение цены за 24 часа
                                        'bidPrice': mathematical.number_round(self.pair['bidPrice']), # bidPrice - ближайшая цена заявки на покупку монеты в стакане
                                        'askPrice': mathematical.number_round(self.pair['askPrice']), # askPrice - ближайшая цена заявки на продажу монеты в стакане
                                        'averagePrice': '0', # averagePrice - усреднённая цена
                                        'trailingPrice': '0', # trailingPrice - цена срабатывания стоп-лосса при трейлинге
                                        'buyPrice': '0', # buyPrice - цена покупки/усреднения ордера
                                        'sellPrice': '0', # sellPrice - цена, по которой выставлен ордер на продажу
                                        'allQuantity': '0', # allQuantity - общее количество купленных монет по данной паре
                                        'freeQuantity': '0', # freeQuantity - свободный баланс монеты
                                        'lockQuantity': '0', # lockQuantity - заблокированный в ордерах на покупку баланс
                                        'orderId': '0', # orderId - id размещённого ордера
                                        'profit': '0', # profit - прибыль по данной паре
                                        'totalQuote': '0', # totalQuote - количество актива в монете
                                        'stepAveraging': '0', # stepAveraging - процент шага последующего усреднения
                                        'numAveraging': '0', # numAveraging - количество усреднений
                                        'statusOrder': 'NO_ORDER'} # statusOrder - состояние ордера по данной паре
                                    for self.key in self.dict_values:
                                        self._checked = db.read('symbols', "WHERE pair = '{}'".format(self._symbol['symbol']))
                                        self._checked = self._checked[0] if len(self._checked) > 0 else {self.key: None}
                                        if self._checked.get(self.key) == None or self._symbol['symbol'] not in self._checked['pair']:
                                            db.write('insert', 'symbols', self.key, self.dict_values[self.key]) if self._checked['pair'] == None else db.write('update', 'symbols', self.key, self.dict_values[self.key], pair = self._checked['pair'])
                                    self.check_keys += 1
                                    cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('Добавлено новых торговых пар: {} ({})'.format(str(self.check_keys), self.pair['symbol']), 'cyan') + '\033[K', end = '\r', flush = True)
                                    break

    def check_symbols(self):
        """Проверяем пары в БД"""
        self.trade_keys = len([item for item in db.read('symbols')])
        self.check_keys = 0
        self.offline_trade = 0
        self.bot_orders = [pair['symbol'] for pair in self.all_orders if len(db.read('symbols', "WHERE pair = '{}'".format(pair['symbol']))) > 0 and pair['symbol'] in db.read('symbols', "WHERE pair = '{}'".format(pair['symbol']), keys = ['pair'])[0]['pair'] and ((self.user_order) or 'xbot_' in pair['clientOrderId'] or db.read('symbols', "WHERE pair = '{}'".format(pair['symbol']), keys = ['statusOrder'])[0]['statusOrder'] == 'USER_BUY_ORDER')]
        self.bot_orders_base_asset_list = [db.read('symbols', "WHERE pair = '{}'".format(_), keys = ['baseAsset'])[0]['baseAsset'] for _ in self.bot_orders]
        self.sell_orders = len(self.bot_orders)
        db.write('updates', 'trade_info', '', '', sell_open_orders = self.sell_orders)
        for symbol in db.read('symbols'):
            #if symbol['pair'] == 'XTZUSDT':
            #    test = 1
            if (symbol['statusOrder'] == 'NO_ORDER' or float(symbol['averagePrice']) == 0 or float(symbol['buyPrice']) == 0 or float(symbol['sellPrice']) == 0) and symbol['pair'] in self.bot_orders:
                logging.warning('Предстартовый пересчёт ошибочных данных по {}: {}'.format(symbol['pair'], symbol['statusOrder']))
                for _ in self.all_orders: #этому масиву нужно сделать rekey по symbol и переписать условие.
                    #if _['symbol'] == 'XTZUSDT':
                    #    test = 1
                    try:
                        if ((self.user_order ) or 'xbot_' in _['clientOrderId']) and _['symbol'] == symbol['pair'] and _['side'] == 'SELL':
                            time.sleep(1)
                            last_orders = main.client.get_all_orders(symbol = symbol['pair'])
                            orders_id = sorted([trades for trades in last_orders], key = lambda trades: trades['orderId'], reverse = True)
                            time.sleep(1)
                            last_trades = main.client.get_my_trades(symbol = symbol['pair'])
                            trades_id = sorted([trades for trades in last_trades], key = lambda trades: trades['orderId'], reverse = True)
                            new_average_buy_price = new_all_quantity = all_quantity = totalQuote = step_average = num_average = 0
                            first_buy = True
                            lastorderid = 0
                            for buy_order in orders_id: # Пересчёт данных об ордере на продажу в случае, если он отсутствует в БД
                                for order_price in trades_id: # Находим цену ордера
                                    if buy_order['orderId'] == order_price['orderId'] and order_price['orderId'] != lastorderid:
                                        lastorderid = buy_order['orderId']
                                        if buy_order['side'] == 'BUY' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == True:
                                            buy_price = order_price['price']
                                            totalQuote += float(buy_order['cummulativeQuoteQty'])
                                            all_quantity += float(mathematical.number_round(buy_order['origQty']))
                                            step_average += step_average + float(self.step_aver) if self.num_aver == True else 0
                                            first_buy = False
                                            num_average += 1
                                        elif buy_order['side'] == 'BUY' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == False:
                                            new_average_buy_price += ((float(buy_price) * all_quantity) + (float(order_price['price']) * float(buy_order['origQty']))) if num_average == 1 else float(order_price['price']) * float(buy_order['origQty'])
                                            new_all_quantity += all_quantity + float(buy_order['origQty']) if num_average == 1 else float(buy_order['origQty'])
                                            totalQuote += float(buy_order['cummulativeQuoteQty'])
                                            step_average += float(self.step_aver) if self.num_aver == True else 0
                                            num_average += 1
                                        elif buy_order['side'] == 'SELL' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == False:
                                            first_buy = 'END'
                                            break
                            average_buy_price = ('{:.%sf}' % mathematical.get_count(symbol['tickSize'])).format(new_average_buy_price / new_all_quantity) if new_average_buy_price != 0 and new_all_quantity != 0 else buy_price
                            db.write('updates', 'symbols', 'pair', symbol['pair'],
                                averagePrice = mathematical.number_round(average_buy_price),
                                buyPrice = mathematical.number_round(buy_price),
                                sellPrice = mathematical.number_round(_['price']),
                                trailingPrice = average_buy_price,
                                allQuantity = mathematical.number_round(_['origQty']),
                                freeQuantity = 0,
                                lockQuantity = 0,
                                orderId = _['orderId'],
                                totalQuote = mathematical.number_round('{:.8f}'.format(totalQuote)),
                                stepAveraging = mathematical.number_round('{:.2f}'.format(step_average)),
                                numAveraging = num_average,
                                statusOrder = 'SELL_ORDER')
                            logging.warning('averagePrice: ' + str(average_buy_price) + ' *** ' + 'allQuantity: ' + str(_['origQty']) + ' *** ' + ' totalQuote: ' + mathematical.number_round('{:.8f}'.format(totalQuote)) + ' *** ' + 'step_average: ' + mathematical.number_round('{:.2f}'.format(step_average)) + ' *** ' + 'numAveraging: ' + mathematical.number_round(num_average))
                            break
                    except Exception as e:
                        logging.error('bot.check_symbols(NO_ORDER):\nexcept: {}\n'.format(str(e)))
            elif symbol['statusOrder'] == 'NO_ORDER' and symbol['baseAsset'] in self.free_coins:
                try:
                    if self.user_order:
                        time.sleep(1)
                        last_orders = main.client.get_all_orders(symbol = symbol['pair'])
                        orders_id = sorted([trades for trades in last_orders], key = lambda trades: trades['orderId'], reverse = True)
                        time.sleep(1)
                        last_trades = main.client.get_my_trades(symbol = symbol['pair'])
                        trades_id = sorted([trades for trades in last_trades], key = lambda trades: trades['orderId'], reverse = True)
                        new_average_buy_price = new_all_quantity = all_quantity = totalQuote = step_average = num_average = 0
                        first_buy = True
                        lastorderid = 0
                        for buy_order in orders_id: # Пересчёт данных об ордере на продажу в случае, если он отсутствует в БД
                            for order_price in trades_id: # Находим цену ордера
                                if buy_order['orderId'] == order_price['orderId'] and order_price['orderId'] != lastorderid and 'arbbot' not in buy_order['clientOrderId']:
                                    lastorderid = buy_order['orderId']
                                    if buy_order['side'] == 'BUY' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == True:
                                        buy_price = order_price['price']
                                        totalQuote += float(buy_order['cummulativeQuoteQty'])
                                        all_quantity += float(mathematical.number_round(buy_order['origQty']))
                                        step_average += step_average + float(self.step_aver) if self.num_aver == True else 0
                                        first_buy = False
                                        num_average += 1
                                    elif buy_order['side'] == 'BUY' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == False:
                                        new_average_buy_price += ((float(buy_price) * all_quantity) + (float(order_price['price']) * float(buy_order['origQty']))) if num_average == 1 else float(order_price['price']) * float(buy_order['origQty'])
                                        new_all_quantity += all_quantity + float(buy_order['origQty']) if num_average == 1 else float(buy_order['origQty'])
                                        totalQuote += float(buy_order['cummulativeQuoteQty'])
                                        step_average += float(self.step_aver) if self.num_aver == True else 0
                                        num_average += 1
                                    elif buy_order['side'] == 'SELL' and buy_order['status'] == 'FILLED' and ('xbot_' in buy_order['clientOrderId'] or self.user_order) and first_buy == False:
                                        first_buy = 'END'
                                        break
                        new_all_quantity = all_quantity if new_all_quantity == 0 else new_all_quantity
                        average_buy_price = ('{:.%sf}' % mathematical.get_count(symbol['tickSize'])).format(new_average_buy_price / new_all_quantity) if new_average_buy_price != 0 and new_all_quantity != 0 else buy_price
                        average_sell_price = ('{:.%sf}' % mathematical.get_count(symbol['tickSize'])).format(float(average_buy_price) + ((float(average_buy_price) / 100) * (float(bot.sell_up) + float(0)))) # Находим новую цену продажи без дельты
                        new_all_quantity = ('{:.%sf}' % mathematical.get_count(symbol['stepSize'])).format(new_all_quantity)
                        self.offline_trade += 1
                        cprint('|OFF-LINE| ' + colored('Исполнено во время бездействия:\033[K', 'cyan')) if self.offline_trade == 1 else ''
                        cprint('|OFF-LINE| ' + colored('FREE_COIN ', 'magenta') + '(MARKED): ' + colored('[' + mathematical.number_round(new_all_quantity) + ']', 'grey', 'on_white') + ' ' + 
                            colored(symbol['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + colored(mathematical.number_round(average_sell_price) + ' ' + symbol['quoteAsset'] + '\033[K', 'yellow'))
                        db.write('updates', 'symbols', 'pair', symbol['pair'],
                            averagePrice = mathematical.number_round(average_buy_price),
                            buyPrice = mathematical.number_round(buy_price),
                            sellPrice = mathematical.number_round(average_sell_price),
                            trailingPrice = mathematical.number_round(average_buy_price),
                            allQuantity = 0,
                            freeQuantity = mathematical.number_round(new_all_quantity),
                            lockQuantity = 0,
                            orderId = 0,
                            totalQuote = mathematical.number_round('{:.8f}'.format(totalQuote)),
                            stepAveraging = mathematical.number_round('{:.2f}'.format(step_average)),
                            numAveraging = num_average,
                            statusOrder = 'FREE_AVERAGING_ORDER')
                        logging.warning('averagePrice: ' + str(average_buy_price) + ' *** ' + 'sellPrice: ' + str(average_sell_price) + ' *** ' + 'freeQuantity: ' + str(new_all_quantity) + ' *** ' + 'allQuantity: ' + str(0) + ' *** ' + ' totalQuote: ' + mathematical.number_round('{:.8f}'.format(totalQuote)) + ' *** ' + 'step_average: ' + mathematical.number_round('{:.2f}'.format(step_average)) + ' *** ' + 'numAveraging: ' + mathematical.number_round(num_average))
                except Exception as e:
                        logging.error('bot.check_symbols(NO_ORDER):\nexcept: {}\n'.format(str(e)))
            elif symbol['statusOrder'] == 'SELL_ORDER' and symbol['pair'] not in self.bot_orders:
                self.offline_trade += 1
                cprint('|OFF-LINE| ' + colored('Исполнено во время бездействия:\033[K', 'cyan')) if self.offline_trade == 1 else ''
                profit = '{:.8f}'.format(float(symbol['sellPrice']) * float(symbol['allQuantity']) - float(symbol['totalQuote']))
                cprint('|OFF-LINE| ' + colored('SELL ', 'red') + '(FILLED): ' + colored('[' + symbol['allQuantity'] + ']', 'grey', 'on_white') + ' ' + 
                    colored(symbol['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + colored(symbol['sellPrice'] + ' ' + symbol['quoteAsset'] + '\033[K', 'yellow'))
                telegram.sell(symbol, profit)
                db.write('updates', 'symbols', 'pair', symbol['pair'], profit = mathematical.number_round('{:.8f}'.format(float(profit) + float(symbol['profit']))))
                db.write('updates', 'trade_info', '', '',
                    sell_filled_orders = int(db.read('trade_info', keys = ['sell_filled_orders'])[0]['sell_filled_orders']) + 1,
                    sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1)
                db.write('insert', 'daily_profit', '', '',
                    day = self.today,
                    quote = symbol['quoteAsset'],
                    profit = profit)
                bot.write_no_order(symbol['pair'])

            elif symbol['statusOrder'] == 'USER_CLOSE_SELL_ORDER':
                self.offline_trade += 1
                cprint('|OFF-LINE| ' + colored('Исполнено во время бездействия:\033[K', 'cyan')) if self.offline_trade == 1 else ''
                profit = symbol['totalQuote']
                cprint('|OFF-LINE| ' + colored('SELL ', 'red') + '(FILLED): ' + colored('[' + symbol['allQuantity'] + ']', 'grey', 'on_white') + ' ' + 
                    colored(symbol['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + colored(symbol['sellPrice'] + ' ' + symbol['quoteAsset'] + '\033[K', 'yellow'))
                telegram.sell(symbol, profit)
                db.write('updates', 'symbols', 'pair', symbol['pair'], profit = mathematical.number_round('{:.8f}'.format(float(profit) + float(symbol['profit']))))
                db.write('updates', 'trade_info', '', '',
                    sell_filled_orders = int(db.read('trade_info', keys = ['sell_filled_orders'])[0]['sell_filled_orders']) + 1,
                    sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1)
                db.write('insert', 'daily_profit', '', '',
                    day = self.today,
                    quote = symbol['quoteAsset'],
                    profit = profit)
                bot.write_no_order(symbol['pair'])

            elif symbol['statusOrder'] == 'TRAILING_STOP_ORDER':
                for balance in self.balances:
                    if balance['asset'] == symbol['baseAsset']:
                        if float(balance['free']) < float(symbol['freeQuantity']) and float(balance['free']) > 0:
                            db.write('updates', 'symbols', 'pair', symbol['pair'],
                                freeQuantity = mathematical.number_round(balance['free']),
                                totalQuote = mathematical.number_round('{:.8f}'.format(float(balance['free']) / float(symbol['freeQuantity']) * float(symbol['totalQuote']))))
                        elif float(balance['free']) == 0:
                            bot.write_no_order(symbol['pair'])
                            db.write('updates', 'symbols', 'pair', symbol['pair'], freeQuantity = mathematical.number_round(balance['free']))

            elif (symbol['statusOrder'] == 'BUY_ORDER' and symbol['pair'] not in self.bot_orders):
                try:
                    last_id = sorted([trades for trades in main.client.get_all_orders(symbol = symbol['pair'])], key = lambda trades: trades['orderId'], reverse = True)
                    find_order = False
                    for buy_order in last_id:
                        if buy_order['side'] == 'BUY' and buy_order['status'] == 'FILLED' and int(buy_order['orderId']) == int(symbol['orderId']):
                            self.offline_trade += 1
                            cprint('|OFF-LINE| ' + colored('Исполнено во время бездействия:\033[K', 'cyan')) if self.offline_trade == 1 else ''
                            db.write('updates', 'symbols', 'pair', symbol['pair'],
                                buyPrice = mathematical.number_round(buy_order['price']),
                                allQuantity = 0,
                                freeQuantity = mathematical.number_round(buy_order['origQty']),
                                lockQuantity = 0,
                                totalQuote = '{:.8f}'.format(float(buy_order['origQty']) * float(buy_order['price'])),
                                stepAveraging = mathematical.number_round('{:.2f}'.format(float(symbol['stepAveraging']) + float(self.step_aver))) if self.num_aver == True else symbol['stepAveraging'],
                                numAveraging = int(symbol['numAveraging']) + 1,
                                statusOrder = 'FREE_SELL_ORDER')
                            find_order = True
                    if find_order == False:
                        bot.write_no_order(symbol['pair'])
                except Exception as e:
                    logging.error('bot.check_symbols(BUY_ORDER):\nexcept: {}\n'.format(str(e)))

            self.check_keys += 1
            cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('Проверено торговых пар: {}/{}'.format(str(self.check_keys), str(self.trade_keys)), 'cyan') + '\033[K', end = '\r', flush = True)
        if self.offline_trade > 0:
            cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('Всего исполнено: {}'.format(str(self.offline_trade)) + '\033[K', 'cyan'))

    def start_sockets(self):
        """Запуск сокетов"""
        bnb_burn_true = True if main.client.toggle_bnb_burn_spot_margin(spotBNBBurn = 'true')['spotBNBBurn'] == True else False
        bnb_burn_color = colored('Включено\033[K', 'green') if bnb_burn_true == True else colored('Ошибка включения\033[K\n', 'red')
        cprint(colored('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ', 'white') + colored('Оплата комиссии биржи в BNB: ', 'cyan') + bnb_burn_color, end = '\r', flush = True) # Включаем оплату комиссии в BNB
        if bnb_burn_true == True:
            exit_trade = threading.Thread(target = bot.exit_sockets, daemon = True)
            exit_trade.start()
            stop_trades = threading.Thread(target = bot.stop_trades, daemon = True)
            stop_trades.start()
            self.binance_socket_manager = BinanceSocketManager(main.client, user_timeout = 60)
            self.conn_start_user_socket = self.binance_socket_manager.start_user_socket(stream.start_user_socket)
            self.conn_search_pair = self.binance_socket_manager.start_ticker_socket(stream.search_pair)
            self.binance_socket_manager.start()
            #self.binance_socket_manager.join()
            #exit_trade.join()
            #stop_trades.join()
            bot.day_profit = len(db.read('daily_profit', condition = "WHERE day NOT LIKE '{}'".format(bot.today), keys = ['id', 'day', 'quote', 'profit']))

    def close_sockets(self):
        """Останавливаем бота и выходим в главное меню"""
        self.binance_socket_manager.close()
        del self.binance_socket_manager, self.conn_start_user_socket, self.conn_search_pair
        time.sleep(2)
        main.print_menu()

    def canceled_order(self, _symbol, _side, _orderId, _clientId):
        """Отмена ордера"""
        try:
            return main.client.cancel_order(symbol = _symbol, orderId = _orderId, newClientOrderId = _clientId)
        except Exception as e:
            if 'code' in e.__dict__:
                if str(e.__dict__['code']) == '-2011':
                    try:
                        order = main.client.get_open_orders(symbol = _symbol)
                        search_order = False
                        for _ in order:
                            if _['side'] == _side and _['clientOrderId'] == _clientId:
                                search_order = True
                                new_canceled_order = main.client.cancel_order(symbol = _symbol, orderId = _['orderId'], newClientOrderId = _clientId)
                                db.write('updates', 'symbols', 'pair', _symbol,
                                    orderId = new_canceled_order['orderId'])
                                return new_canceled_order
                        if search_order == False:
                            for pair_none in db.read('symbols', "WHERE pair = '{}'".format(_symbol)):
                                if pair_none['statusOrder'] == 'SELL_ORDER':
                                    profit = mathematical.number_round('{:.8f}'.format(float(pair_none['sellPrice']) * float(pair_none['allQuantity']) - float(pair_none['totalQuote'])))
                                    telegram.sell(pair_none, profit)
                            bot.write_no_order(_symbol)
                    except Exception as e:
                        logging.error('bot.canceled_order():\n_symbol: {}\n_side: {}\n_orderId: {}\n_clientId: {}\nexcept: {}\n'.format(_symbol, _side, _orderId, _clientId, str(e)))

    def write_no_order(self, pair):
        """Обнуление пары в БД из-за ошибки"""
        db.write('updates', 'symbols', 'pair', pair,
            averagePrice = 0,
            buyPrice = 0,
            sellPrice = 0,
            trailingPrice = 0,
            allQuantity = 0,
            freeQuantity = 0,
            lockQuantity = 0,
            orderId = 0,
            totalQuote = 0,
            stepAveraging = 0,
            numAveraging = 0,
            statusOrder = 'NO_ORDER')

    def bnb_buy(self):
        """Докупка BNB"""
        if float(self.quote_balances['BNB']['free']) < float(self.min_bnb):
            main.exchange = main.client.get_exchange_info()
            main.tickers = main.client.get_ticker()
            for pair in main.tickers:
                for quote in self.quote_balances:
                    if pair['symbol'] == 'BNB' + quote:
                        for step in main.exchange['symbols']:
                            if step['symbol'] == pair['symbol']:
                                try:
                                    order_quantity = ('{:.%sf}' % mathematical.get_count(mathematical.number_round(step['filters'][2]['stepSize']))).format((float(mathematical.number_round(step['filters'][3]['minNotional']))) / float(pair['askPrice']))
                                    while float(order_quantity) * float(pair['askPrice']) <= float(step['filters'][3]['minNotional']):
                                        order_quantity = ('{:.%sf}' % mathematical.get_count(mathematical.number_round(step['filters'][2]['stepSize']))).format(float(order_quantity) + float(step['filters'][2]['stepSize']))
                                    if float(self.quote_balances[quote]['free']) > float(order_quantity) * float(pair['askPrice']):
                                        bnb_order = main.client.order_market_buy(symbol = pair['symbol'], quantity = order_quantity) # Пытаемся докупить BNB
                                        self.quote_balances['BNB']['free'] = '{:.8f}'.format(float(self.quote_balances['BNB']['free']) + float(bnb_order['origQty']))
                                        self.quote_balances[quote]['free'] = '{:.8f}'.format(float(self.quote_balances[quote]['free']) - float(bnb_order['cummulativeQuoteQty']))
                                        return True
                                except Exception as e:
                                    logging.error('bot.filter():\nexcept: {}\n'.format(str(e)))
                                    return False
        else:
            return True

    def bnb_comission(self, N, Y, n):
        """Подсчёт комиссии и объёма ордера"""
        db.write('insert', 'bnb_burn', '', '',
            day = bot.today,
            pair = N,
            size = Y,
            comission = n)

    def exit_sockets(self):
        """Реагируем на команду -e"""
        while True:
            self.key = input('')
            if self.key == '-e': # and self.conn_start_user_socket != None:
                break
        bot.close_sockets()

    def stop_trades(self):
        """Проверяем работу бота по количеству биржевых тиков"""
        self.ticks = 0
        while True:
            self.ticks += 1
            if self.ticks > 60:
                telegram.stop_bot()
                break
            time.sleep(1)

class Telegram(): # Telegram

    def __init__(self):
        """Запускаем Telegram-бота и количество попыток отправить оповещение"""
        self.bot = telebot.TeleBot(bot.tg_token)
        self.try_send = 0

    def statistics_update(self):
        """Обновление торговой статистики в Telegram"""
        if self.try_send <= 3 and bot.tg_notification == True:
            try:
                _oa = _fb = ''
                for quote in bot.quote_balances:
                    if quote in bot.quote_asset_list:
                        _oa = _oa + '\n' + mathematical.number_round(bot.quote_balances[quote]['overall']) + ' ' + quote
                        _fb = _fb + '\n' + mathematical.number_round(bot.quote_balances[quote]['free']) + ' ' + quote
                _text = ('💼 Общий баланс: {}\n🔑 Свободный баланс: {}\n📂 Открытые ордера: {}\n🔥 Успешных сделок: {}\n🎯 Дельта: {}%\n⏳ Обновлено в {}').format(_oa, _fb, str(len([1 for _ in db.read('symbols', condition = "WHERE statusOrder NOT LIKE 'NO_ORDER%'")])), str(int(float(db.read('trade_info')[0]['sell_filled_orders']))), mathematical.number_round('{:.2f}'.format(sum([float(_['priceChangePercent']) for _ in db.read('symbols')]) / len([1 for _ in db.read('symbols')]))), str(datetime.datetime.now().strftime('%H:%M:%S'))) # Текст для описания в Telegram-канале
                self.bot.set_chat_description(bot.tg_name, _text) # Обновление статистики в Telegram-канале
            except Exception as e:
                logging.error('telegram.statistics_update():\nexcept: {}\n'.format(str(e)))
                self.try_send += 1
                telegram.statistics_update()
        else:
            self.try_send = 0

    def daily_profit(self, daily_profit):
        """Дневная статистика c отправкой в Telegram"""
        if self.try_send <= 3:
            try:
                main.exchange = main.client.get_exchange_info()
                main.tickers = main.client.get_ticker()
                bnb_burn = db.read('bnb_burn', condition = "WHERE day NOT LIKE '{}'".format(bot.today), keys = ['id', 'day', 'pair', 'size', 'comission'])
                average_percent = db.read('average_percent', condition = "WHERE day NOT LIKE '{}'".format(bot.today), keys = ['id', 'day', 'percent'])
                _daily_text = ''
                _day_list = list()
                _quote_list = list()
                _quote_profit = dict()
                _bnb_burn = dict()
                for item in daily_profit:
                    if item['day'] not in _day_list:
                        _day_list.append(item['day'])
                    if item['quote'] not in _quote_list:
                        _quote_list.append(item['quote'])
                for _day in _day_list:
                    if bot.tg_notification == True:
                        for _quote in _quote_list:
                            _sells = 0
                            _profit = 0
                            for _item in daily_profit:
                                if _item['quote'] == _quote:
                                    _sells += 1
                                    _profit += float(_item['profit'])
                            _daily_text = _daily_text + '\n   {}'.format(_quote + (' ' * (6 - len(_quote)))) + ' | ' + mathematical.number_round('{:.8f}'.format(_profit))
                            _quote_profit.update({_quote: str(_profit)})
                        for _fiat_currency in bot.fiat_currencies_list:
                            _profit = 0
                            for _quote in _quote_profit:
                                for _info in main.exchange['symbols']:
                                    if _info['quoteAsset'] == _fiat_currency and _info['baseAsset'] == _quote:
                                        for _fiat in main.tickers:
                                            if _fiat['symbol'] == _info['symbol']:
                                                _profit += float(_quote_profit[_quote]) * float(_fiat['lastPrice'])
                            _daily_text = _daily_text + '\n   {}'.format(_fiat_currency + (' ' * (6 - len(_fiat_currency)))) + ' | ' + mathematical.number_round('{:.2f}'.format(_profit)) if _profit != 0 else _daily_text
                        for item in bnb_burn:
                            if item['pair'] not in _bnb_burn:
                                _bnb_burn.update({item['pair']: {'size': item['size'], 'comission': item['comission']}})
                            else:
                                _bnb_burn[item['pair']]['comission'] = '{:.8f}'.format(float(_bnb_burn[item['pair']]['comission']) + float(item['comission']))
                                _bnb_burn[item['pair']]['size'] = '{:.8f}'.format(float(_bnb_burn[item['pair']]['size']) + float(item['size']))
                        _daily_size = '\n🔹 Объём:'
                        _daily_comission = '\n🔸 Комиссия:'
                        for item in _bnb_burn:
                            _daily_comission = _daily_comission + '\n   {} {}'.format(mathematical.number_round(_bnb_burn[item]['comission']), item)
                            _daily_size = _daily_size + '\n   {} {}'.format(mathematical.number_round(_bnb_burn[item]['size']), item)
                        _daily_percent = 0
                        for item in average_percent:
                            _daily_percent += float(item['percent'])
                        _avr_daily_profit = 0
                        if _daily_percent > 0 and len(average_percent) > 0:
                            _avr_daily_profit = _daily_percent / len(average_percent)
                        _daily_percent = '\n🔺 Средняя прибыль: {}%'.format(mathematical.number_round('{:.2f}'.format(_avr_daily_profit)))
                        _text = ('<code>🔔 Статистика за {}\n✅ Сделок закрыто: {}\n   Валюта | Прибыль{}{}{}{}</code>').format(_day, _sells, _daily_text, _daily_percent, _daily_size, _daily_comission)
                        _tg_message = self.bot.send_message(chat_id = bot.tg_name, text = _text, parse_mode = 'HTML')
                        self.bot.pin_chat_message(chat_id = bot.tg_name, message_id = _tg_message.message_id)
                    db.write('delete', 'bnb_burn', 'day', _day)
                    db.write('delete', 'daily_profit', 'day', _day)
                    db.write('delete', 'average_percent', 'day', _day)
                    self.try_send = bot.day_profit = 0
            except Exception as e:
                logging.error('telegram.daily_profit():\nexcept: {}\n'.format(str(e)))
                self.try_send += 1
                telegram.daily_profit(db.read('daily_profit', condition = "WHERE day NOT LIKE '{}'".format(bot.today), keys = ['id', 'day', 'quote', 'profit']))
        else:
            self.try_send = 0

    def sell(self, symbol, profit):
        """Отправляем сообщение о продаже"""
        if self.try_send <= 3:
            try:
                _trailing_list = db.read('trailing_orders', condition = "WHERE pair = '{}'".format(symbol['pair']))
                if len(_trailing_list) == 0:
                    _text = ('<code>📝 {}/{}\n📉 Средняя цена покупки: {}\n📈 Цена продажи: {}\n💵 Объём: {}\n💎 Прибыль: {}% ({} {})</code>').format(symbol['baseAsset'], symbol['quoteAsset'], symbol['averagePrice'], symbol['sellPrice'], symbol['allQuantity'], mathematical.number_round('{:.2f}'.format((float(symbol['sellPrice']) / float(symbol['averagePrice']) - 1) * 100)), mathematical.number_round(profit), symbol['quoteAsset'])
                else:
                    _space = sorted([space for space in _trailing_list], key = lambda space: len(space['p']), reverse = True)
                    _q = 0
                    _avg = 0
                    _word = '   Цена '
                    _len = len(_word) if len(_word) >= len(_space[0]['p']) else len(_space[0]['p'])
                    _orders_dict = 'Всего ордеров: {}\n{}{}| Объём'.format(str(len(_trailing_list)), _word, (' ' * (_len - len(_word) + 4)))
                    for trailing_sales in _trailing_list:
                        _q += float(trailing_sales['q'])
                        _avg += float(trailing_sales['p']) * float(trailing_sales['q'])
                        _orders_dict = '{}\n   {}{}| {}'.format(_orders_dict, trailing_sales['p'], (' ' * (_len - len(trailing_sales['p']) + 1)), trailing_sales['q'])
                    _avg = mathematical.number_round(('{:.%sf}' % mathematical.get_count(symbol['tickSize'])).format(_avg / _q))
                    _q = mathematical.number_round(('{:.%sf}' % mathematical.get_count(symbol['stepSize'])).format(_q))
                    _orders_dict = '📈 Средняя цена продажи: {}\n📊 {}'.format(mathematical.number_round(('{:.%sf}' % mathematical.get_count(symbol['tickSize'])).format(float(_avg))), _orders_dict)
                    _all_quantity = mathematical.number_round(('{:.%sf}' % mathematical.get_count(symbol['stepSize'])).format(float(_q)))
                    _percent = mathematical.number_round('{:.2f}'.format((float(_avg) / float(symbol['averagePrice']) - 1) * 100))
                    _profit = mathematical.number_round('{:.8f}'.format((float(_avg) * float(_q)) - (float(symbol['averagePrice']) * float(_q))))
                    _text = ('<code>📝 {}/{}\n📉 Средняя цена покупки: {}\n{}\n💵 Общий объём: {}\n💎 Прибыль: {}% ({} {})</code>').format(
                        symbol['baseAsset'],
                        symbol['quoteAsset'],
                        mathematical.number_round(symbol['averagePrice']),
                        _orders_dict,
                        _all_quantity,
                        _percent,
                        _profit,
                        symbol['quoteAsset'])
                if bot.tg_notification == True:
                    logging.info('chat_id: {}\ntext: {}\n'.format(bot.tg_name, _text))
                    self.bot.send_message(chat_id = bot.tg_name, text = _text, parse_mode = 'HTML')
                    self.try_send = 0
            except Exception as e:
                logging.error('telegram.sell():\nexcept: {}\n'.format(str(e)))
                self.try_send += 1
                telegram.sell(symbol, profit)
        else:
            self.try_send = 0
        db.write('insert', 'average_percent', '', '', day = bot.today, percent = (mathematical.number_round('{:.2f}'.format((float(symbol['sellPrice']) / float(symbol['averagePrice']) - 1) * 100)) if len(_trailing_list) == 0 else _percent))
        db.write('delete', 'trailing_orders', 'pair', symbol['pair'])

    def stop_bot(self): # Отправляем в Telegram пуш об ошибке бота
        if bot.tg_notification == True:
            if self.try_send <= 3:
                try:
                    _text = '<code>❗ Стрим веб-сокета остановлен!\n🔌 Проверьте бота на работоспособность!</code>'
                    _tg_message = self.bot.send_message(chat_id = bot.tg_name, text = _text, parse_mode = 'HTML')
                    self.bot.pin_chat_message(chat_id = bot.tg_name, message_id = _tg_message.message_id)
                except Exception as e:
                    logging.error('telegram.stop_bot():\nexcept: {}\n'.format(str(e)))
                    self.try_send += 1
                    telegram.stop_bot()
            else:
                self.try_send = 0

class Stream():
    """Класс стримов  веб-сокеты"""
    def __init__(self):
        self.timer_new_day = bot.new_day() # Сколько секунд осталось до конца дня
        self.telegram_statistics_timer = 86400 # Таймер стрима
        self.trade_list = list() # Список торгующихся пар
        self.best_coin_list = dict() # Словарь приоритетных монет
        self.delta_dict = dict() # Словарь всех дельт котируемых активов
        self.spinner_delta = 0
        self.event = False

    def search_pair(self, msg_sp):
        """Ивентовый стрим реагируем на события"""
        bot.ticks = 0
        now_time = time.time()
        timer_daily_profit = bot.new_day()
        bnb_buy = bot.bnb_buy()

        if bot.timestamp + 3000 <= now_time: # keepalive и перезапуск бота в случае ошибки
            try:
                main.client.stream_keepalive(bot.conn_start_user_socket)
                bot.timestamp = now_time
            except:
                bot.binance_socket_manager.close()
                del bot.binance_socket_manager, bot.conn_start_user_socket, bot.conn_search_pair
                bot.connect()

        if timer_daily_profit > self.timer_new_day or bot.day_profit > 0: # Пуш суточной статистики в Telegram
            bot.today = datetime.date.today()
            daily_profit = db.read('daily_profit', condition = "WHERE day NOT LIKE '{}'".format(bot.today), keys = ['id', 'day', 'quote', 'profit'])
            if len(daily_profit) > 0:
                telegram.daily_profit(daily_profit)
            self.timer_new_day = bot.new_day()

        if now_time - bot.stream_timer >= 0 and bnb_buy == True: # Поиск подходящих по словиям пар для торговли
            bot.stream_timer = now_time
            for quote in bot.quote_asset_list:
                quantity_sell_order = 0
                if self.event == False and quote != 'BUSD':
                    overall_balance = 0 # Общий баланс котируемого актива
                    averaging_value = 0 # Величина усреднений всех активов из таблицы symbols
                    open_orders = 0 # Количество открытых ордеров к котируемому активу
                    daily_delta = 0 # Суточная дельта
                    self.trade_pairs = sorted([_ for _ in db.read('symbols') if _['quoteAsset'] == quote], key = lambda _: ((float(_['askPrice']) - float(_['sellPrice'])) / (float(_['sellPrice']) / 100)) if float(_['askPrice']) != 0 and float(_['sellPrice']) != 0 and _['statusOrder'] == 'SELL_ORDER' else float(_['priceChangePercent']), reverse = False) # Сортируем все пары из белого списка по значению суточного падения или падения askPrice от SellPrice
                    for _ in self.trade_pairs:
                        self.best_coin_list[_['pair']] = {'baseAsset': _['baseAsset'], 'quoteAsset': _['quoteAsset'], 'askPrice': _['askPrice'], 'sellPrice': _['sellPrice']}
                        self.trade_list.append(_['pair']) if _['pair'] not in self.trade_list else None
                        overall_balance = '{:.8f}'.format(float(overall_balance) + float(_['totalQuote']))
                        averaging_value += float(_['stepAveraging']) if float(_['stepAveraging']) >= 1 else float(_['numAveraging'])
                        open_orders += 1 if _['statusOrder'] != 'NO_ORDER' else 0
                        daily_delta += float(_['priceChangePercent']) # Находим % среднесуточного показателя цен всех пар за 24 часа
                    overall_balance = '{:.8f}'.format(float(overall_balance) + float(bot.quote_balances[quote]['free']))
                    bot.quote_balances[quote]['overall'] = overall_balance
                    averaging_value = float(mathematical.number_round('{:.2f}'.format(averaging_value)))
                    len_trade_pairs = len(self.trade_pairs)
                    self.print_daily_delta = float(mathematical.number_round('{:.2f}'.format(daily_delta / len_trade_pairs))) if daily_delta != 0 and len_trade_pairs > 0 else 0 # Значение дельты для cprint
                    self.daily_delta = float('{:.3f}'.format(self.print_daily_delta / 5)) if float(bot.sell_up) * 3 >= self.print_daily_delta >= float(bot.buy_down) * 3 and self.print_daily_delta != 0 and bot.delta_percent == True else 0 if bot.delta_percent == False else float('{:.3f}'.format(float(bot.sell_up) * 3 / 5)) if float(bot.sell_up) * 3 <= self.print_daily_delta else float('{:.3f}'.format(float(bot.buy_down) * 3 / 5)) # Применяем 3/5 дельту для бота, если соответствуют условия
                    self.delta_dict[quote] = {'delta': str(self.print_daily_delta)}
                    max_trade_pairs = int(float(bot.max_trade_pairs) * (float(bot.quote_balances[quote]['free']) / float(overall_balance)) * (open_orders / averaging_value * averaging_value / open_orders)) if bot.auto_trade_pairs == True and int(bot.max_trade_pairs) != -1 and float(bot.quote_balances[quote]['free']) != 0 and float(overall_balance) != 0 and open_orders != 0 and averaging_value != 0 else int(bot.max_trade_pairs) # Автокорректировка ботом максимального количества открытых позиций
                    max_trade_pairs = max_trade_pairs if int(bot.max_trade_pairs) >= max_trade_pairs >= 0 or max_trade_pairs == -1 or bot.auto_trade_pairs == False else int(bot.max_trade_pairs) if max_trade_pairs > int(bot.max_trade_pairs) else 0
                    bot.quote_balances[quote]['max_trade_pairs'] = str(max_trade_pairs)
                    symbols_dict = dict()
                    for pair_sp in msg_sp:
                        symbols_dict.update({pair_sp['s'] : {'priceChangePercent': pair_sp['P'],'bidPrice': pair_sp['b'], 'askPrice': pair_sp['a']}}) if pair_sp['s'] in self.trade_list else None
                        if pair_sp['s'] == (quote + 'USDT'):
                            bot.quote_balances[quote]['equivalent'] = pair_sp['a']
                    db.write('updates_sp', 'symbols', '', '', _dict = symbols_dict)
                    self.trade_pairs = sorted([_ for _ in db.read('symbols') if _['quoteAsset'] == quote], key = lambda _: ((float(_['askPrice']) - float(_['sellPrice'])) / (float(_['sellPrice']) / 100)) if float(_['askPrice']) != 0 and float(_['sellPrice']) != 0 and _['statusOrder'] == 'SELL_ORDER' else float(_['priceChangePercent']), reverse = False)
                    for pair_db in self.trade_pairs:
                        step_size = str(mathematical.get_count(pair_db['stepSize']))
                        tick_size = str(mathematical.get_count(pair_db['tickSize']))
                        ncoi = 'xbot_' + pair_db['baseAsset'] + pair_db['quoteAsset']

                        if pair_db['statusOrder'] == 'CANCELED_BUY_ORDER': # Продажа возможных свободных монет и обнуление данных по вручную отменённому ордеру на покупку
                            try:
                                main.client.order_market_sell( # Пытаемся распродать свободные монеты от незавершённых сделок
                                    symbol = pair_db['pair'],
                                    quantity = pair_db['freeQuantity'],
                                    newClientOrderId = ncoi)
                                self.event = True
                                break
                            except Exception as e:
                                bot.write_no_order(pair_db['pair'])
                                logging.error('stream.search_pair(SELL CANCELED_BUY_ORDE):\n{}\nexcept: {}\n'.format(str(pair_db), str(e)))

                        elif pair_db['statusOrder'] == 'CANCELED_SELL_ORDER': # Расчёт и усреднение отменённого ботом ордера на продажу
                            try:
                                free_quantity = pair_db['allQuantity'] # Размер отменённого ордера
                                average_quantity = mathematical.number_round(('{:.%sf}' % step_size).format(float(free_quantity) * (float(bot.quantity_aver) + (float(pair_db['numAveraging']) / 10)))) # Размер нового ордера для усреднения
                                main.client.order_market_buy( # Открываем позицию по активу
                                    symbol = pair_db['pair'],
                                    quantity = average_quantity,
                                    newClientOrderId = ncoi)
                                self.event = True
                                break
                            except Exception as e:
                                db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                            allQuantity = 0,
                                            freeQuantity = pair_db['allQuantity'],
                                            lockQuantity = 0,
                                            statusOrder = 'FREE_AVERAGING_ORDER')# fixed allQuantity to freeQuantity and allQuantity = 0 and lockQuantity = 0
                                logging.error('stream.search_pair(BUY CANCELED_SELL_ORDER):\n{}\nexcept: {}\n'.format(str(pair_db), str(e)))

                        elif pair_db['statusOrder'] == 'FREE_AVERAGING_ORDER': # Размещение ордера на продажу после усреднения позиции
                            try:
                                main.client.order_limit_sell( # Размещаем ордер на продажу усреднённого пользователем ордера
                                    symbol = pair_db['pair'],
                                    quantity = pair_db['freeQuantity'],
                                    price = pair_db['sellPrice'],
                                    newClientOrderId = ncoi)
                                self.event = True
                                break
                            except Exception as e:
                                logging.error('stream.search_pair(SELL FREE_AVERAGING_ORDER):\n{}\nexcept: {}\n'.format(str(pair_db), str(e)))

                        elif pair_db['statusOrder'] == 'USER_AVERAGING_ORDER': # Обновляем данные о позиции после полностью исполненного пользовательского ордера на усреднение
                            try:
                                canceled_order = bot.canceled_order( # Отменяем ордер на продажу для объединения с ручным усредняющим ордером
                                    pair_db['pair'],
                                    'SELL',
                                    pair_db['orderId'],
                                    ncoi)
                                if canceled_order != None:
                                    average_buy_price = ('{:.%sf}' % tick_size).format(((float(pair_db['averagePrice']) * float(canceled_order['origQty'])) + (float(pair_db['buyPrice']) * float(pair_db['freeQuantity']))) / (float(canceled_order['origQty']) + float(pair_db['freeQuantity']))) # Находим цену усреднения
                                    average_sell_price = ('{:.%sf}' % tick_size).format(float(average_buy_price) + ((float(average_buy_price) / 100) * (float(bot.sell_up) + float(self.daily_delta)))) # Находим новую цену продажи
                                    db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                        averagePrice = average_buy_price,
                                        sellPrice = average_sell_price,
                                        trailingPrice = average_buy_price,
                                        freeQuantity = ('{:.%sf}' % step_size).format(float(canceled_order['origQty']) + float(pair_db['freeQuantity'])),
                                        statusOrder = 'FREE_AVERAGING_ORDER')
                                    self.event = True
                                    break
                                else:
                                    bot.write_no_order(pair_db['pair'])
                            except Exception as e:
                                logging.error('stream.search_pair(CANCELED USER_AVERAGING_ORDER):\n{}\nexcept: {}\n'.format(str(pair_db), str(e)))
                                bot.write_no_order(pair_db['pair'])

                        elif pair_db['statusOrder'] == 'FREE_SELL_ORDER': # Размещаем ордер на продажу после полностью исполненного ордера на первую покупку монеты
                            try:
                                _price = ('{:.%sf}' % tick_size).format(float(pair_db['averagePrice']) + (float(pair_db['averagePrice']) / 100 * (float(bot.sell_up) + self.daily_delta)))
                                _price = ('{:.%sf}' % tick_size).format(float(_price) + float(pair_db['tickSize'])) if float(_price) == float(pair_db['averagePrice']) else _price
                                main.client.order_limit_sell( # Размещаем ордер на продажу
                                    symbol = pair_db['pair'],
                                    quantity = pair_db['freeQuantity'],
                                    price = _price if int(pair_db['numAveraging']) > 1 else pair_db['sellPrice'] if float(pair_db['sellPrice']) > 0 else ('{:.%sf}' % tick_size).format(float(pair_db['buyPrice']) * (1 + (float(bot.sell_up) + self.daily_delta) / 100)),
                                    newClientOrderId = ncoi)
                                self.event = True
                                break
                            except Exception as e:
                                logging.error('stream.search_pair(SELL FREE_SELL_ORDER):\n{}\n_price: {}\nexcept: {}\n'.format(str(pair_db), _price, str(e)))

                        elif pair_db['statusOrder'] == 'SELL_ORDER': # Отменяем текущий ордер для дальнейшего усреднения или активируем трейлинг-стоп при удовлетворяющих условиях
                            try:
                                quantity_sell_order += 1
                                if bot.trailing_stop == True: # Отмена ордера для активации трейлинг-стопа
                                    if float(pair_db['askPrice']) >= float(pair_db['sellPrice']) * (1 - (float(bot.trailing_percent) * (float(bot.sell_up) + (self.daily_delta if self.daily_delta > 0 else 0)) / float(bot.sell_up) / 100)) and float(pair_db['askPrice']) * float(bot.quote_balances[quote]['equivalent']) >= float(bot.trailing_price):
                                        db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                            statusOrder = 'ACTIVATE_TRAILING_STOP_ORDER')
                                if int(bot.max_aver) > int(pair_db['numAveraging']): # Отмена ордера для усреднения позиции
                                    if float(pair_db['askPrice']) <= float(('{:.%sf}' % tick_size).format(float(pair_db['buyPrice']) - (float(pair_db['buyPrice']) / 100 * (float(pair_db['stepAveraging']) - float(bot.buy_down) + self.daily_delta)))):
                                        order_quantity = ('{:.%sf}' % step_size).format(float(pair_db['allQuantity']) * (float(bot.quantity_aver) + (float(pair_db['stepAveraging']) / 10)))
                                        if float(bot.quote_balances[quote]['free']) - (float(order_quantity) * float(pair_db['askPrice'])) >= float(bot.quote_balances[quote]['overall']) / 100 * float(bot.min_balance):
                                            bot.canceled_order( # Отменяем ордер для усреднения
                                                pair_db['pair'],
                                                'SELL',
                                                pair_db['orderId'],
                                                ncoi)
                                            self.event = True
                                            break
                            except Exception as e:
                                logging.error('stream.search_pair(SELL_ORDER):\n{}\norder_quantity: {}\nexcept: {}\n'.format(str(pair_db), order_quantity, str(e)))

                        elif pair_db['statusOrder'] == 'ACTIVATE_TRAILING_STOP_ORDER': # Отменяем ордер на продажу для активации трейлинг-стопа
                            try:
                                bot.canceled_order( # Отменяем ордер для активации трейлинга
                                    pair_db['pair'],
                                    'SELL',
                                    pair_db['orderId'],
                                    ncoi)
                                self.event = True
                                break
                            except Exception as e:
                                logging.error('stream.search_pair(CANCELED ACTIVATE_TRAILING_STOP_ORDER):\n{}\nexcept: {}\n'.format(str(pair_db), str(e)))

                        elif pair_db['statusOrder'] == 'TRAILING_STOP_ORDER': # Частичная или полная продажа по трейлинг-стопу
                            try:
                                if float(pair_db['askPrice']) >= float(pair_db['sellPrice']): # передвигаем триггер трейлинг-стопа вверх при росте цены
                                    type_trailing_order = 'NEW_PRICE'
                                    new_sell_price = pair_db['sellPrice']
                                    new_trailing_price = pair_db['trailingPrice']
                                    while float(pair_db['askPrice']) > float(new_sell_price) or float(new_sell_price) == float(pair_db['sellPrice']):
                                        new_sell_price = mathematical.number_round(('{:.%sf}' % tick_size).format(float(new_sell_price) + float(pair_db['tickSize'])))
                                        new_trailing_price = mathematical.number_round(('{:.%sf}' % tick_size).format(float(new_trailing_price) + float(pair_db['tickSize'])))
                                    db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                        sellPrice = new_sell_price,
                                        trailingPrice = new_trailing_price)
                                elif float(pair_db['askPrice']) < float(pair_db['sellPrice']) * (1 - (float(bot.trailing_percent) * (float(bot.sell_up) + self.daily_delta) / float(bot.sell_up) / 100)): # При падении цены на n% от триггера
                                    type_trailing_order = 'NEW_QUANTITY'
                                    trailing_quantity = mathematical.number_round(('{:.%sf}' % step_size).format(float(pair_db['minNotional']) if float(bot.trailing_part) == 0 else float(pair_db['freeQuantity']) / 100 * float(bot.trailing_part)))
                                    while float(pair_db['minNotional']) > float(trailing_quantity) * float(pair_db['bidPrice']):
                                        trailing_quantity = mathematical.number_round(('{:.%sf}' % step_size).format(float(trailing_quantity) + float(pair_db['stepSize'])))
                                    if float(pair_db['freeQuantity']) - float(trailing_quantity) < float(trailing_quantity) * (1 - float(bot.trailing_percent)) and ((float(pair_db['freeQuantity']) - float(trailing_quantity)) * float(pair_db['averagePrice']) <= float(pair_db['minNotional']) or float(pair_db['bidPrice']) <= float(pair_db['trailingPrice'])) and float(pair_db['bidPrice']) >= float(pair_db['averagePrice']): # Закрытие позиции
                                        type_trailing_order = 'FULL'
                                        main.client.order_market_sell(symbol = pair_db['pair'], quantity = pair_db['freeQuantity'], newClientOrderId = ncoi) # Создаём ордер на продажу (закрыть позицию)
                                        self.event = True
                                        break
                                    elif (float(pair_db['freeQuantity']) - float(trailing_quantity)) * float(pair_db['trailingPrice']) > float(pair_db['minNotional']) and float(pair_db['bidPrice']) * 1.0015 > float(pair_db['trailingPrice']): # Частичная продажа
                                        type_trailing_order = 'PARTIAL'
                                        ncoim = 'xbottrailing_' + pair_db['baseAsset'] + pair_db['quoteAsset']
                                        main.client.order_market_sell(symbol = pair_db['pair'], quantity = trailing_quantity, newClientOrderId = ncoim) # Создаём ордер на продажу (частичная продажа)
                                        self.event = True
                                        break
                                    elif float(pair_db['askPrice']) <= float(pair_db['averagePrice']) * (1 - float(bot.trailing_percent)) and float(pair_db['freeQuantity']) * float(pair_db['sellPrice']) > float(pair_db['minNotional']): # Выставляем обратно ордер на продажу
                                        type_trailing_order = 'SELL'
                                        db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                            statusOrder = 'FREE_SELL_ORDER')
                                    elif float(('{:.%sf}' % tick_size).format(float(pair_db['buyPrice']) - (float(pair_db['buyPrice']) / 100 * (float(pair_db['stepAveraging']) - float(bot.buy_down) - self.daily_delta)))): # Усреднение позиции при упущенном трейлинг-стопе
                                        type_trailing_order = 'AVERAGING'
                                        db.write('updates', 'symbols', 'pair', pair_db['pair'],
                                            statusOrder = 'FREE_AVERAGING_ORDER')
                            except Exception as e:
                                logging.error('stream.search_pair({} TRAILING_STOP_ORDER):\n{}\nexcept: {}\n'.format(type_trailing_order, str(pair_db), str(e)))

                        elif pair_db['statusOrder'] == 'NO_ORDER': # Осуществляем первую покупку монеты при удовлетворяющих условиях и отсутствии ордера по ней
                            try:
                                if max_trade_pairs > open_orders >= 0 or max_trade_pairs == -1:
                                    if float(pair_db['askPrice']) * float(bot.quote_balances[quote]['equivalent']) >= float(bot.min_price) and float(pair_db['priceChangePercent']) <= float(bot.daily_percent):
                                        if (pair_db['baseAsset'] not in bot.bot_orders_base_asset_list and bot.double_asset == False) or bot.double_asset == True:
                                            order_quantity = float(('{:.%sf}' % step_size).format(float(bot.min_order) * float(pair_db['minNotional']) / float(pair_db['askPrice'])))
                                            while order_quantity * float(pair_db['askPrice']) < float(pair_db['minNotional']):
                                                order_quantity += float(pair_db['stepSize'])
                                            if float(bot.quote_balances[quote]['free']) - (order_quantity * float(pair_db['askPrice'])) >= float(bot.quote_balances[quote]['overall']) / 100 * float(bot.min_balance):
                                                order_quantity = ('{:.%sf}' % step_size).format(order_quantity)
                                                main.client.order_market_buy( # Открываем позицию по активу
                                                    symbol = pair_db['pair'],
                                                    quantity = order_quantity,
                                                    newClientOrderId = ncoi)
                                                self.event = True
                                                break
                            except Exception as e:
                                logging.error('stream.search_pair(BUY NO_ORDER):\n{}\norder_quantity: {}\nexcept: {}\n'.format(str(pair_db), order_quantity, str(e)))

            self.print_daily_delta = '{:.2f}'.format(sum([float(self.delta_dict[_]['delta']) for _ in self.delta_dict]) / len(self.delta_dict))
            spinner_delta_color = 'red' if self.spinner_delta > float(self.print_daily_delta) else 'green' if self.spinner_delta != float(self.print_daily_delta) else 'yellow'
            row_delta = '↓' if self.spinner_delta > float(self.print_daily_delta) else '↑' if self.spinner_delta < float(self.print_daily_delta) else ''
            self.spinner_delta = float(self.print_daily_delta)
            sync_time = [time.time() - float(_['E']) / 1000 for _ in msg_sp]
            sync_time = sum(sync_time) / len(sync_time)
            sync_times = colored('•', 'green' if sync_time <= 3 else 'yellow' if sync_time <= 10 else 'red')
            print_max_trade_pairs = ' '
            print_balances = ' '
            best_coin = sorted([_ for _ in self.best_coin_list if float(self.best_coin_list[_]['askPrice']) != 0 and float(self.best_coin_list[_]['sellPrice']) != 0], key = lambda _: float(self.best_coin_list[_]['sellPrice']) / float(self.best_coin_list[_]['askPrice']), reverse = False)
            print_best_coin = '' if len(best_coin) == 0 else colored('Л: ', 'cyan') + colored('{} {}% ({} {})'.format(self.best_coin_list[best_coin[0]]['baseAsset'], mathematical.number_round('{:.2f}'.format(((float(self.best_coin_list[best_coin[0]]['sellPrice']) / float(self.best_coin_list[best_coin[0]]['askPrice'])) - 1) * 100)), self.best_coin_list[best_coin[0]]['sellPrice'], self.best_coin_list[best_coin[0]]['quoteAsset']), 'yellow') + ' | '
            for quote in bot.quote_balances:
                print_max_trade_pairs = print_max_trade_pairs + bot.quote_balances[quote]['max_trade_pairs'] + ' ' + quote + ' * ' if quote in bot.quote_asset_list else print_max_trade_pairs + ''
                print_balances = print_balances + mathematical.number_round(bot.quote_balances[quote]['free']) + ' ' + quote + ' * ' if quote in bot.quote_asset_list else print_balances + ''
            cprint('|{}| {} {} {}{} | {}{}\033[K'.format(str(datetime.datetime.now().strftime('%H:%M:%S')), colored('{:.2f}'.format(time.time() - now_time) + 's', 'cyan'), sync_times, (colored('М:', 'cyan') + colored(print_max_trade_pairs[:-3], 'yellow') + ' | ' if bot.auto_trade_pairs == True else ''), colored('Б:' + colored(print_balances[:-3], 'yellow'), 'cyan'), print_best_coin, colored('Д: ' + colored(mathematical.number_round(self.print_daily_delta) + '% ', spinner_delta_color), 'cyan') + colored(row_delta, spinner_delta_color)), 'white', end = '\r', flush = True)

        elif bnb_buy == False:
            cprint('Не удалось купить BNB!\033[K', 'red', end = '\r', flush = True)

        if self.telegram_statistics_timer - timer_daily_profit > 60: # Обновление торговой статистики в Telegram-канале
            telegram.statistics_update()
            self.telegram_statistics_timer = timer_daily_profit

        bot.stream_timer = time.time()
        stream.event = False

    def start_user_socket(self, msg_sus):
        """Стрим аккаунта"""

        if msg_sus['e'] == 'outboundAccountPosition':
            for quote in bot.quote_balances:
                for coin in msg_sus.get('B'):
                    if coin['a'] == quote:
                        bot.quote_balances[quote]['free'] = coin['f']

        if msg_sus['e'] == 'executionReport':
            s = msg_sus.get('s') # Символ
            S = msg_sus.get('S') # BUY|SELL
            p = msg_sus.get('p') # Цена
            q = msg_sus.get('q') # Общее количество монет в ордере
            l = msg_sus.get('l') # Последнее выполненное количество монет в ордере
            L = msg_sus.get('L') # Цена последнего исполнения
            z = msg_sus.get('z') # Суммарное заполненное количество монет в ордере
            Z = msg_sus.get('Z') # Суммарная стоимость ордера
            Y = msg_sus.get('Y') # Частичная стоимость ордера
            i = msg_sus.get('i') # ID ордера
            X = msg_sus.get('X') # Статус ордера
            o = msg_sus.get('o') # Тип ордера
            c = msg_sus.get('c') # ID клиента
            C = msg_sus.get('C') # ID клиента (отменённый)
            n = msg_sus.get('n') # Количество комиссии
            N = msg_sus.get('N') # Актив комиссии

            if (X == 'FILLED' or X == 'PARTIALLY_FILLED') and ('xbot_' in c or 'xbottrailing_' in c):
                bot.bnb_comission(N, Y, n)

            if 'xbottrailing_' in c and X == 'FILLED' and S == 'SELL':
                order.trailing(s, q, L, Y, Z) # Рыночный трейлинг-ордер

            elif 'xbot_' in c or 'xbot_' in C:
                if X == 'NEW': # Новый ордер
                    order.new(s, S, p, L, q, i, X, o)
                elif X == 'PARTIALLY_FILLED': # Часть ордера исполнилась
                    order.partially_filled(s, S, q, l, z, Y, i, o)
                elif X == 'FILLED': # Ордер исполнился
                    order.filled(s, S, p, q, L, l, z, Z, Y, i, o)
                elif X == 'CANCELED': # Ордер отменён ботом
                    order.cancel(s, S, p, L, q, z, i, X, c, C, o)

            elif 'xbot_' not in c and bot.user_order == True:
                if X == 'NEW': # Новый ордер
                    order.new(s, S, p, L, q, i, X, o, C)
                elif X == 'PARTIALLY_FILLED': # Часть ордера исполнилась
                    order.partially_filled(s, S, q, l, z, Y, i, o)
                elif X == 'FILLED': # Новый ордер
                    order.filled(s, S, p, q, L, l, z, Z, Y, i, o)

class Orders():
    """Работа с ордерами"""

    def trailing(self, s, q, L, Y, Z):
        """Частичная продажа по рынку при трейлинге"""
        for pair_db in db.read('symbols', "WHERE pair = '{}'".format(s)):
            if pair_db['statusOrder'] == 'TRAILING_STOP_ORDER':
                profit = mathematical.number_round('{:.8f}'.format(float(Y) - float(pair_db['totalQuote'])))
                db.write('insert', 'trailing_orders', '', '',
                    pair = s,
                    p = mathematical.number_round(L),
                    q = mathematical.number_round(q))
                db.write('updates', 'symbols', 'pair', s,
                    sellPrice = mathematical.number_round(('{:.%sf}' % mathematical.get_count(pair_db['tickSize'])).format(float(L) - float(pair_db['tickSize']))),
                    freeQuantity = mathematical.number_round(('{:.%sf}' % mathematical.get_count(pair_db['stepSize'])).format(float(pair_db['freeQuantity']) - float(q))),
                    totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) - (float(Z))))
                db.write('insert', 'daily_profit', '', '',
                    day = bot.today,
                    quote = pair_db['quoteAsset'],
                    profit = profit if float(profit) > 0 else 0)
                cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('SELL ', 'magenta') + '(MARKET): ' + 
                    colored('[' + mathematical.number_round(q) + ']', 'grey', 'on_white') + ' ' + colored(pair_db['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + 
                    colored(mathematical.number_round(L) + ' ' + pair_db['quoteAsset'] + '\033[K', 'yellow'))
                break

    def new(self, s, S, p, L, q, i, X, o, C = 'xbot_'):
        """Новый ордер"""
        for pair_db in db.read('symbols', "WHERE pair = '{}'".format(s)):
            _side = 'BUY ' if S == 'BUY' else 'SELL'
            _color = 'green' if S == 'BUY' else 'red'
            _p = mathematical.number_round(p)
            _q = mathematical.number_round(q)
            cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('%s ' % _side, '%s' % _color) + '(PLACED): ' + 
                colored('[' + _q + ']', 'grey', 'on_white') + ' ' + colored(pair_db['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + 
                colored(_p + ' ' + pair_db['quoteAsset'] + '\033[K', 'yellow')) if o != 'MARKET' else None

            if S == 'BUY':
                statusOrder = 'BUY_ORDER' if 'xbot_' in str(C) else 'USER_BUY_ORDER' if pair_db['statusOrder'] == 'NO_ORDER' else 'USER_AVERAGE_ORDER'
                
                if pair_db['statusOrder'] == 'NO_ORDER':
                    bot.bot_orders_base_asset_list.append(pair_db['baseAsset']) if pair_db['baseAsset'] not in bot.bot_orders_base_asset_list else None
                    db.write('updates', 'symbols', 'pair', s,
                        averagePrice = _p if o != 'MARKET' else pair_db['averagePrice'],
                        trailingPrice = _p if o != 'MARKET' else pair_db['trailingPrice'],
                        buyPrice = _p if o != 'MARKET' else pair_db['buyPrice'],
                        lockQuantity = _q,
                        orderId = i,
                        statusOrder = statusOrder)

                elif pair_db['statusOrder'] == 'SELL_ORDER':
                    db.write('updates', 'symbols', 'pair', s,
                        lockQuantity = _q,
                        orderId = i if o == 'MARKET' else pair_db['orderId'],
                        statusOrder = statusOrder if o == 'MARKET' else pair_db['statusOrder'])

                elif pair_db['statusOrder'] == 'CANCELED_SELL_ORDER':
                    db.write('updates', 'symbols', 'pair', s,
                        lockQuantity = _q,
                        orderId = i,
                        statusOrder = 'AVERAGE_SELL_ORDER')

            elif S == 'SELL' and pair_db['statusOrder'] != 'TRAILING_STOP_ORDER':
                db.write('updates', 'trade_info', '', '',
                    sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) + 1)

                if pair_db['statusOrder'] == 'FREE_SELL_ORDER' or pair_db['statusOrder'] == 'CANCELED_BUY_ORDER':
                    db.write('updates', 'symbols', 'pair', s,
                        sellPrice = _p,
                        allQuantity = _q,
                        freeQuantity = 0,
                        lockQuantity = 0,
                        orderId = i,
                        statusOrder = 'SELL_ORDER')

                elif pair_db['statusOrder'] == 'FREE_AVERAGING_ORDER':
                    db.write('updates', 'symbols', 'pair', s,
                        sellPrice = _p,
                        allQuantity = _q,
                        freeQuantity = 0,
                        lockQuantity = 0,
                        orderId = i,
                        statusOrder = 'SELL_ORDER')

            break

    def partially_filled (self, s, S, q, l, z, Y, i, o):
        """Частично исполненный ордер"""
        for pair_db in db.read('symbols', "WHERE pair = '{}'".format(s)):
            if int(pair_db['orderId']) == int(i) or (o == 'LIMIT' and pair_db['statusOrder'] == 'SELL_ORDER' and S == 'BUY'):
                _step_size = mathematical.get_count(pair_db['stepSize'])
                _qz = mathematical.number_round(('{:.%sf}' % _step_size).format(float(q) - float(z)))

                if S == 'BUY':
                    db.write('updates', 'symbols', 'pair', s,
                        freeQuantity = mathematical.number_round(z),
                        lockQuantity = _qz,
                        totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) + float(Y)))

                elif S == 'SELL':
                    db.write('updates', 'symbols', 'pair', s,
                        allQuantity = _qz if o == 'LIMIT' else pair_db['allQuantity'],
                        freeQuantity = _qz if o == 'MARKET' else pair_db['freeQuantity'],
                        totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) - float(Y)))
                break

    def filled(self, s, S, p, q, L, l, z, Z, Y, i, o):
        """Исполненный ордер"""
        for pair_db in db.read('symbols', "WHERE pair = '{}'".format(s)):
            if int(pair_db['orderId']) == int(i) or (o == 'MARKET' and pair_db['statusOrder'] == 'TRAILING_STOP_ORDER') or (o == 'LIMIT' and pair_db['statusOrder'] == 'SELL_ORDER' and S == 'BUY'):
                _side = 'BUY ' if S == 'BUY' else 'SELL'
                _color = 'green' if S == 'BUY' else 'red' if pair_db['statusOrder'] != 'TRAILING_STOP_ORDER' else 'magenta'
                _step_size = mathematical.get_count(pair_db['stepSize'])
                _tick_size = mathematical.get_count(pair_db['tickSize'])
                _p = mathematical.number_round(L) if float(L) != 0 else mathematical.number_round(p)
                _q = mathematical.number_round(('{:.%sf}' % _step_size).format(float(l) + float(pair_db['freeQuantity'])))
                cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('%s ' % _side, '%s' % _color) + '(FILLED): ' + 
                    colored('[' + mathematical.number_round(z) + ']', 'grey', 'on_white') + ' ' + colored(pair_db['baseAsset'], 'magenta', attrs=['bold']) + ' for ' + 
                    colored(_p + ' ' + pair_db['quoteAsset'] + '\033[K', 'yellow'))

                if S == 'BUY':

                    if pair_db['statusOrder'] == 'BUY_ORDER':
                        db.write('updates', 'symbols', 'pair', s,
                            buyPrice = _p,
                            averagePrice = _p,
                            trailingPrice = _p,
                            freeQuantity = _q,
                            lockQuantity = 0,
                            stepAveraging = mathematical.number_round('{:.2f}'.format(float(pair_db['stepAveraging']) + float(bot.step_aver))),
                            numAveraging = int(pair_db['stepAveraging']) + 1,
                            totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) + float(Y)),
                            statusOrder = 'FREE_SELL_ORDER')

                    elif pair_db['statusOrder'] == 'USER_BUY_ORDER':
                        db.write('updates', 'symbols', 'pair', s,
                            buyPrice = _p,
                            averagePrice = _p,
                            trailingPrice = _p,
                            freeQuantity = _q,
                            lockQuantity = 0,
                            totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) + float(Y)),
                            statusOrder = 'FREE_SELL_ORDER')

                    elif pair_db['statusOrder'] == 'USER_AVERAGE_ORDER' or pair_db['statusOrder'] == 'SELL_ORDER':
                        db.write('updates', 'symbols', 'pair', s,
                            buyPrice = _p,
                            freeQuantity = _q,
                            lockQuantity = 0,
                            totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) + float(Y)),
                            statusOrder = 'USER_AVERAGING_ORDER')

                    elif pair_db['statusOrder'] == 'AVERAGE_SELL_ORDER':
                        _average_buy_price = ('{:.%sf}' % _tick_size).format(((float(pair_db['averagePrice']) * float(pair_db['allQuantity'])) + (float(_p) * float(_q))) / (float(pair_db['allQuantity']) + float(_q))) # Находим цену усреднения
                        _average_sell_price = ('{:.%sf}' % _tick_size).format(float(_average_buy_price) + ((float(_average_buy_price) / 100) * (float(bot.sell_up) + float(stream.daily_delta)))) # Находим новую цену продажи
                        db.write('updates', 'symbols', 'pair', pair_db['pair'],
                            buyPrice = _p,
                            averagePrice = _average_buy_price,
                            sellPrice = _average_sell_price,
                            trailingPrice = _average_buy_price,
                            allQuantity = 0,
                            freeQuantity = ('{:.%sf}' % _step_size).format(float(_q) + float(pair_db['allQuantity'])),
                            lockQuantity = 0,
                            stepAveraging = mathematical.number_round('{:.2f}'.format(float(pair_db['stepAveraging']) + float(bot.step_aver))),
                            numAveraging = int(pair_db['numAveraging']) + 1,
                            totalQuote = '{:.8f}'.format(float(pair_db['totalQuote']) + float(Y)),
                            statusOrder = 'FREE_AVERAGING_ORDER')

                if S == 'SELL':
                    logging.info('s Символ: {}\n S BUY|SELL: {}\n p Цена: {}\n q Общее количество монет в ордере: {}\n L Цена последнего исполнения: {}\n l Последнее выполненное количество монет в ордере: {}\n z Суммарное заполненное количество монет в ордере: {}\n Z Суммарная стоимость ордера: {}\n Y Частичная стоимость ордера: {}\n i ID ордера: {}\n o Тип ордера: {}\n'''.format(s, S, p, q, L, l, z, Z, Y, i, o))
                    bot.bot_orders_base_asset_list.remove(pair_db['baseAsset']) if pair_db['baseAsset'] in bot.bot_orders_base_asset_list else None

                    if pair_db['statusOrder'] == 'SELL_ORDER':
                        profit = mathematical.number_round('{:.8f}'.format(float(Z) - float(pair_db['totalQuote'])))
                        bot.write_no_order(s)
                        db.write('updates', 'symbols', 'pair', s, profit = mathematical.number_round('{:.8f}'.format(float(pair_db['profit']) + float(profit))))
                        db.write('insert', 'daily_profit', '', '',
                            day = bot.today,
                            quote = pair_db['quoteAsset'],
                            profit = profit if float(profit) > 0 else 0)
                        logging.info('pair_db: {}\nprofit: {}\n'.format(pair_db, profit))
                        telegram.sell(pair_db, profit)

                    if pair_db['statusOrder'] == 'TRAILING_STOP_ORDER':
                        profit = mathematical.number_round('{:.8f}'.format(float(Y) - float(pair_db['totalQuote'])))
                        bot.write_no_order(s)
                        db.write('updates', 'symbols', 'pair', s, profit = mathematical.number_round('{:.8f}'.format(float(pair_db['profit']) + float(profit))))
                        db.write('insert', 'daily_profit', '', '',
                            day = bot.today,
                            quote = pair_db['quoteAsset'],
                            profit = profit if float(profit) > 0 else 0)
                        db.write('insert', 'trailing_orders', '', '',
                            pair = s,
                            p = mathematical.number_round(L),
                            q = mathematical.number_round(q))
                        telegram.sell(pair_db, profit)

                    elif pair_db['statusOrder'] == 'CANCELED_BUY_ORDER':
                        bot.write_no_order(s)

                    db.write('updates', 'trade_info', '', '',
                        sell_filled_orders = int(db.read('trade_info', keys = ['sell_filled_orders'])[0]['sell_filled_orders']) + 1,
                        sell_open_orders = int(db.read('trade_info', keys = ['sell_open_orders'])[0]['sell_open_orders']) - 1) if pair_db['statusOrder'] != 'CANCELED_BUY_ORDER' else None
                break

    def cancel(self, s, S, p, L, q, z, i, X, c, C, o):
        """Отменённый ордер"""
        for pair_db in db.read('symbols', "WHERE pair = '{}'".format(s)):
            if int(pair_db['orderId']) == int(i):
                _side = 'BUY ' if S == 'BUY' else 'SELL'
                _color = 'magenta' if pair_db['statusOrder'] == 'ACTIVATE_TRAILING_STOP_ORDER' else 'yellow'
                _p = mathematical.number_round(L) if float(L) != 0 else mathematical.number_round(p)
                cprint('|' + str(datetime.datetime.now().strftime('%H:%M:%S')) + '| ' + colored('%s ' % _side, '%s' % _color) + ('(CANCELED): ') + 
                    colored('[' + mathematical.number_round(q) + ']', 'grey', 'on_white') + ' ' + colored(pair_db['baseAsset'], 'magenta', attrs=['bold']) + ' for ' +
                    colored(_p + ' ' + pair_db['quoteAsset'], 'yellow') + colored(' ID ' + str(i) + '\033[K', 'grey', attrs=['bold']))

                if pair_db['statusOrder'] == 'BUY_ORDER' or pair_db['statusOrder'] == 'USER_BUY_ORDER':
                    if float(pair_db['freeQuantity']) * float(pair_db['askPrice']) > float(pair_db['minNotional']):
                        db.write('updates', 'symbols', 'pair', s,
                            statusOrder = 'CANCELED_BUY_ORDER')
                    else:
                        bot.write_no_order(s)
                        bot.bot_orders_base_asset_list.remove(pair_db['baseAsset']) if pair_db['baseAsset'] in bot.bot_orders_base_asset_list else None

                elif pair_db['statusOrder'] == 'SELL_ORDER' and 'xbot_' not in c:
                    db.write('updates', 'symbols', 'pair', s,
                        allQuantity = 0,
                        freeQuantity = pair_db['allQuantity'],
                        orderId = 0,
                        statusOrder = 'FREE_SELL_ORDER')

                elif pair_db['statusOrder'] == 'SELL_ORDER' and 'xbot_' in c:
                    db.write('updates', 'symbols', 'pair', s,
                        statusOrder = 'CANCELED_SELL_ORDER')

                elif pair_db['statusOrder'] == 'ACTIVATE_TRAILING_STOP_ORDER':
                    db.write('updates', 'symbols', 'pair', pair_db['pair'],
                        allQuantity = 0,
                        freeQuantity = mathematical.number_round(q),
                        orderId = 0,
                        statusOrder = 'TRAILING_STOP_ORDER')
                break

            elif int(pair_db['orderId']) != int(i) and (o == 'LIMIT' and pair_db['statusOrder'] == 'SELL_ORDER' and S == 'BUY'):
                if pair_db['statusOrder'] == 'SELL_ORDER':
                    db.write('updates', 'symbols', 'pair', s,
                        lockQuantity = 0,
                        statusOrder = 'USER_AVERAGING_ORDER')

if __name__ == '__main__':
    """Главный блок"""

    if os.path.exists("xbot_log.log"):
        if os.path.getsize("xbot_log.log") > 1000000:
            os.remove("xbot_log.log")
    logging.basicConfig(level = logging.INFO, filename = ('xbot_log.log'), format = '%(asctime)s %(levelname)s: %(message)s', encoding='utf-8') # Включаем логирование
    var = Initialization() # Инициализируем все важные переменные
    mathematical = MathFunc() # Присваиваем mathematical математический класс
    db = DataBase() # Присваиваем database класс работы с БД
    main = Main() # Присваиваем main класс предстартового взаимодействия с ботом
    db.check_table_trade_pairs()
    bot = Bot() # Присваиваем запуск бота
    stream = Stream() # Присваиваем торговые стримы
    telegram = Telegram() # Пуши в Telegram
    order = Orders() # Ордера из веб-сокета
    main.print_menu() # Печатаем меню