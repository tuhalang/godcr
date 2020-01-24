package values

import "image/color"

var (
	//  ProgressBarGray indicates the level of sync progress that is yet to be completed.
	ProgressBarGray = color.RGBA{230, 234, 237, 255}

	// ProgressBarGreen indicates the level of sync progress that has been completed.
	ProgressBarGreen = color.RGBA{65, 190, 83, 255}

	// walletSyncBoxGray is the background color of wallet sync boxes.
	WalletSyncBoxGray = color.RGBA{243, 245, 246, 255}
)