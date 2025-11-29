package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Edge_tts(text string) {
	// 1. 设置你想说的话
	voice := "zh-CN-XiaoxiaoNeural" // 晓晓
	outputFile := "tts.mp3"

	// 获取当前路径，防止找不到文件
	cwd, _ := os.Getwd()
	fullPath := filepath.Join(cwd, outputFile)

	fmt.Println("正在合成语音 (调用 edge-tts)...")

	// 2. 使用 Go 的 exec 库调用命令行工具
	// 相当于在终端执行：edge-tts --text "..." --write-media "output.mp3" --voice "..."
	cmd := exec.Command("edge-tts",
		"--text", text,
		"--write-media", fullPath,
		"--voice", voice,
	)

	// 捕获标准输出和错误，方便调试
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("合成失败: %v\n请确保你已经运行了 'pip install edge-tts'\n", err)
		return
	}

	//fmt.Println("合成成功！文件保存在:", fullPath)

	// 3. 播放音频 (Windows)
	playCmd := exec.Command("cmd", "/c", "start", fullPath)
	playCmd.Run()

	////4.删除文件
	//err = os.Remove(fullPath)
}

func Win_tts(text string) {
	// 直接调用 Windows PowerShell 的说话功能
	psCommand := fmt.Sprintf(`Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s')`, text)

	cmd := exec.Command("powershell", "-Command", psCommand)

	fmt.Println("正在说话...")
	err := cmd.Run()
	if err != nil {
		fmt.Println("出错了:", err)
	}
}
