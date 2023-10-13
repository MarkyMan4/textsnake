package main

import "github.com/charmbracelet/lipgloss"

var baseStyle = lipgloss.NewStyle().
	Width(2).
	Height(1)

var snakeStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Background(lipgloss.Color("#14db8f"))

var pelletStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Bold(true).
	Background(lipgloss.Color("#db1481")).
	AlignHorizontal(lipgloss.Center)

var boardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#bf00ff"))

var gameOverStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a32a2a"))

var scoreStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#14db8f"))

var gameOverText = `
::::::::      :::     ::::    ::::  ::::::::::   ::::::::  :::     ::: :::::::::: :::::::::  
:+:    :+:   :+: :+:   +:+:+: :+:+:+ :+:         :+:    :+: :+:     :+: :+:        :+:    :+: 
+:+         +:+   +:+  +:+ +:+:+ +:+ +:+         +:+    +:+ +:+     +:+ +:+        +:+    +:+ 
:#:        +#++:++#++: +#+  +:+  +#+ +#++:++#    +#+    +:+ +#+     +:+ +#++:++#   +#++:++#:  
+#+   +#+# +#+     +#+ +#+       +#+ +#+         +#+    +#+  +#+   +#+  +#+        +#+    +#+ 
#+#    #+# #+#     #+# #+#       #+# #+#         #+#    #+#   #+#+#+#   #+#        #+#    #+# 
 ########  ###     ### ###       ### ##########   ########      ###     ########## ###    ###
`
